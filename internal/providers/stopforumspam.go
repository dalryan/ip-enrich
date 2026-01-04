package providers

import (
	"encoding/json"
	"fmt"

	"github.com/dalryan/ip-enrich/internal/provider"
)

// StopForumSpamResponse represents the response from Stop Forum Spam API.
type StopForumSpamResponse struct {
	Success int `json:"success"`
	IP      struct {
		Value            string  `json:"value"`
		Frequency        int     `json:"frequency"`
		Appears          int     `json:"appears"`
		Confidence       float64 `json:"confidence"`
		LastSeen         string  `json:"lastseen"`
		DelegatedCountry string  `json:"delegated_country"`
		Country          string  `json:"country"`
		ASN              int     `json:"asn"`
	} `json:"ip"`
}

// StopForumSpam implements the Provider interface for Stop Forum Spam.
type StopForumSpam struct {
	provider.BaseProvider
}

// NewStopForumSpam creates a new Stop Forum Spam provider.
func NewStopForumSpam() *StopForumSpam {
	return &StopForumSpam{
		BaseProvider: provider.BaseProvider{
			ProviderName: "Stop Forum Spam",
			ProviderID:   "stopforumspam",
			URLTemplate:  "https://api.stopforumspam.org/api?json&ip={ip}",
		},
	}
}

// ParseResponse parses the Stop Forum Spam API response into normalized fields.
func (s *StopForumSpam) ParseResponse(body []byte, statusCode int) (*provider.Result, error) {
	if statusCode != 200 {
		return provider.NewErrorResult(s, statusCode, fmt.Errorf("unexpected status code: %d", statusCode)), nil
	}

	var resp StopForumSpamResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if resp.Success != 1 {
		return provider.NewErrorResult(s, statusCode, fmt.Errorf("API returned success=%d", resp.Success)), nil
	}

	result := provider.NewSuccessResult(s, statusCode, resp)
	return result, nil
}

func init() {
	provider.Register(NewStopForumSpam())
}
