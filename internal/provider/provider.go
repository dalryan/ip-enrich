// Package provider defines the interface and types for IP enrichment data sources.
package provider

import (
	"context"
	"net/http"
)

// Provider defines the interface that all IP enrichment sources must implement.
type Provider interface {
	// Name returns a human-readable name (e.g., "Shodan")
	Name() string

	// ID returns a unique identifier (e.g., "shodan")
	ID() string

	// BuildRequest constructs an HTTP request for the given IP address.
	BuildRequest(ctx context.Context, ip string) (*http.Request, error)

	// ParseResponse parses the raw response body into a Result.
	ParseResponse(body []byte, statusCode int) (*Result, error)
}

// Result represents the normalized output from any provider.
type Result struct {
	// ProviderID is the unique identifier of the provider that produced this result
	ProviderID string `json:"provider_id"`

	// ProviderName is the human-readable name of the provider
	ProviderName string `json:"provider_name"`

	// StatusCode is the HTTP status code from the response
	StatusCode int `json:"status_code"`

	// Error contains any error message if Success is false
	Error string `json:"error,omitempty"`

	// Raw contains the original parsed response (provider-specific struct)
	Raw any `json:"raw,omitempty"`
}
