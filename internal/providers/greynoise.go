package providers

import (
	"encoding/json"
	"fmt"

	"github.com/dalryan/ip-enrich/internal/provider"
)

// GreyNoiseResponse represents the response from GreyNoise Community API.
type GreyNoiseResponse struct {
	IP             string `json:"ip"`
	Noise          bool   `json:"noise"`
	Riot           bool   `json:"riot"`
	Classification string `json:"classification"`
	Name           string `json:"name"`
	Link           string `json:"link"`
	LastSeen       string `json:"last_seen"`
	Message        string `json:"message"`
}

// GreyNoise implements the Provider interface for GreyNoise Community API.
type GreyNoise struct {
	provider.BaseProvider
}

// NewGreyNoise creates a new GreyNoise provider.
func NewGreyNoise() *GreyNoise {
	return &GreyNoise{
		BaseProvider: provider.BaseProvider{
			ProviderName: "GreyNoise",
			ProviderID:   "greynoise",
			URLTemplate:  "https://api.greynoise.io/v3/community/{ip}",
		},
	}
}

// ParseResponse parses the GreyNoise API response.
func (g *GreyNoise) ParseResponse(body []byte, statusCode int) (*provider.Result, error) {
	// GreyNoise returns 404 for unknown IPs
	if statusCode == 404 {
		return provider.NewSuccessResult(g, statusCode, nil), nil
	}

	if statusCode != 200 {
		return provider.NewErrorResult(g, statusCode, fmt.Errorf("unexpected status code: %d", statusCode)), nil
	}

	var resp GreyNoiseResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	result := provider.NewSuccessResult(g, statusCode, resp)

	return result, nil
}

func init() {
	provider.Register(NewGreyNoise())
}
