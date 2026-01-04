package providers

import (
	"encoding/json"
	"fmt"

	"github.com/dalryan/ip-enrich/internal/provider"
)

// ShodanResponse represents the response from Shodan's InternetDB API.
type ShodanResponse struct {
	IP        string   `json:"ip"`
	Ports     []int    `json:"ports"`
	Hostnames []string `json:"hostnames"`
	CPEs      []string `json:"cpes"`
	Tags      []string `json:"tags"`
	Vulns     []string `json:"vulns"`
}

// Shodan implements the Provider interface for Shodan's InternetDB.
type Shodan struct {
	provider.BaseProvider
}

// NewShodan creates a new Shodan provider.
func NewShodan() *Shodan {
	return &Shodan{
		BaseProvider: provider.BaseProvider{
			ProviderName: "Shodan",
			ProviderID:   "shodan",
			URLTemplate:  "https://internetdb.shodan.io/{ip}",
		},
	}
}

// ParseResponse parses the Shodan API response into normalized fields.
func (s *Shodan) ParseResponse(body []byte, statusCode int) (*provider.Result, error) {
	// Shodan returns 404 for IPs not in their database
	if statusCode == 404 {
		return provider.NewSuccessResult(s, statusCode, nil), nil
	}

	if statusCode != 200 {
		return provider.NewErrorResult(s, statusCode, fmt.Errorf("unexpected status code: %d", statusCode)), nil
	}

	var resp ShodanResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	result := provider.NewSuccessResult(s, statusCode, resp)

	return result, nil
}

func init() {
	provider.Register(NewShodan())
}
