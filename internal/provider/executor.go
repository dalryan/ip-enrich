package provider

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// MaxBodySize defines the maximum bytes we will read from any provider (5MB).
// This just a light safeguard against memory exhaustion.
const MaxBodySize = 5 * 1024 * 1024

// Executor runs providers concurrently and collects results.
type Executor struct {
	client  *http.Client
	timeout time.Duration
}

// ExecutorOption configures an Executor.
type ExecutorOption func(*Executor)

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) ExecutorOption {
	return func(e *Executor) {
		e.timeout = d
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(c *http.Client) ExecutorOption {
	return func(e *Executor) {
		e.client = c
	}
}

// NewExecutor creates a new provider executor.
func NewExecutor(opts ...ExecutorOption) *Executor {
	e := &Executor{
		timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(e)
	}

	if e.client == nil {
		e.client = &http.Client{
			Timeout: e.timeout,
		}
	}

	return e
}

// ResultCallback is called when a provider completes (success or failure).
type ResultCallback func(result *Result)

// Execute runs all providers concurrently for the given IP.
// The callback is called for each result as it completes.
// Returns all results when complete.
func (e *Executor) Execute(ctx context.Context, ip string, providers []Provider, callback ResultCallback) []*Result {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		results []*Result
	)

	for _, p := range providers {
		wg.Add(1)
		go func(p Provider) {
			defer wg.Done()

			result := e.executeOne(ctx, ip, p)

			mu.Lock()
			results = append(results, result)
			mu.Unlock()

			if callback != nil {
				callback(result)
			}
		}(p)
	}

	wg.Wait()
	return results
}

// executeOne runs a single provider and returns the result.
func (e *Executor) executeOne(ctx context.Context, ip string, p Provider) *Result {
	req, err := p.BuildRequest(ctx, ip)
	if err != nil {
		return NewErrorResult(p, 0, fmt.Errorf("failed to build request: %w", err))
	}

	resp, err := e.client.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.Canceled) {
			return NewErrorResult(p, 0, fmt.Errorf("operation cancelled"))
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return NewErrorResult(p, 0, fmt.Errorf("timeout exceeded"))
		}
		return NewErrorResult(p, 0, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	limitReader := io.LimitReader(resp.Body, MaxBodySize)
	body, err := io.ReadAll(limitReader)
	if err != nil {
		return NewErrorResult(p, resp.StatusCode, fmt.Errorf("read body failed: %w", err))
	}

	result, err := p.ParseResponse(body, resp.StatusCode)
	if err != nil {
		return NewErrorResult(p, resp.StatusCode, err)
	}

	return result
}

// ExecuteAsync starts provider execution and returns a channel of results.
// The channel is closed when all providers complete.
func (e *Executor) ExecuteAsync(ctx context.Context, ip string, providers []Provider) <-chan *Result {
	results := make(chan *Result, len(providers))

	go func() {
		defer close(results)
		e.Execute(ctx, ip, providers, func(r *Result) {
			results <- r
		})
	}()

	return results
}
