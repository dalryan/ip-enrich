package output

import (
	"fmt"
	"io"

	"github.com/dalryan/ip-enrich/internal/provider"
)

// Formatter defines the interface for output formatters.
type Formatter interface {
	Format(report *Report) error
}

// Report contains all data for a single IP enrichment run.
type Report struct {
	IP        string             `json:"ip"`
	Timestamp string             `json:"timestamp"`
	Results   []*provider.Result `json:"results"`
}

// NewReport creates a Report from provider results.
func NewReport(ip, timestamp string, results []*provider.Result) *Report {
	return &Report{
		IP:        ip,
		Timestamp: timestamp,
		Results:   results,
	}
}

// GetFormatter returns a formatter for the given format name.
func GetFormatter(format string, w io.Writer) (Formatter, error) {
	switch format {
	case "json":
		return NewJSONFormatter(w, false), nil
	case "pretty":
		return NewJSONFormatter(w, true), nil
	default:
		return nil, fmt.Errorf("unknown format: %s (supported: json, pretty)", format)
	}
}
