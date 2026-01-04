package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// BaseProvider provides common functionality for providers.
// Embed this in your provider implementation to get sensible defaults.
type BaseProvider struct {
	ProviderName string
	ProviderID   string
	URLTemplate  string
	Headers      map[string]string
	Method       string
}

// Name returns the provider's display name.
func (b *BaseProvider) Name() string {
	return b.ProviderName
}

// ID returns the provider's unique identifier.
func (b *BaseProvider) ID() string {
	return b.ProviderID
}

// BuildRequest creates a basic HTTP request with the IP substituted into the URL template.
// Override this method if you need custom request building (POST body, auth, etc.).
func (b *BaseProvider) BuildRequest(ctx context.Context, ip string) (*http.Request, error) {
	url := strings.ReplaceAll(b.URLTemplate, "{ip}", ip)

	method := b.Method
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	// TODO: make this use the actual version not just a hardcoded string
	req.Header.Set("User-Agent", "dalryan/ip-enrich")

	for k, v := range b.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

// NewErrorResult creates a Result representing an error.
func NewErrorResult(p Provider, statusCode int, err error) *Result {
	return &Result{
		ProviderID:   p.ID(),
		ProviderName: p.Name(),
		StatusCode:   statusCode,
		Error:        err.Error(),
	}
}

// NewSuccessResult creates a successful Result with the raw data.
func NewSuccessResult(p Provider, statusCode int, raw any) *Result {
	return &Result{
		ProviderID:   p.ID(),
		ProviderName: p.Name(),
		StatusCode:   statusCode,
		Raw:          raw,
	}
}
