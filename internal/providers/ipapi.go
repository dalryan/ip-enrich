package providers

import (
	"encoding/json"
	"fmt"

	"github.com/dalryan/ip-enrich/internal/provider"
)

// IPAPIResponse represents the response from ipapi.is API.
type IPAPIResponse struct {
	IP           string `json:"ip"`
	RIR          string `json:"rir"`
	IsBogon      bool   `json:"is_bogon"`
	IsDatacenter bool   `json:"is_datacenter"`
	IsTor        bool   `json:"is_tor"`
	IsProxy      bool   `json:"is_proxy"`
	IsVPN        bool   `json:"is_vpn"`
	IsAbuser     bool   `json:"is_abuser"`
	Company      struct {
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Type   string `json:"type"`
	} `json:"company"`
	ASN struct {
		ASN    int    `json:"asn"`
		Org    string `json:"org"`
		Route  string `json:"route"`
		Domain string `json:"domain"`
		Type   string `json:"type"`
	} `json:"asn"`
	Location struct {
		Continent     string  `json:"continent"`
		Country       string  `json:"country"`
		CountryCode   string  `json:"country_code"`
		State         string  `json:"state"`
		City          string  `json:"city"`
		Latitude      float64 `json:"latitude"`
		Longitude     float64 `json:"longitude"`
		Zip           string  `json:"zip"`
		Timezone      string  `json:"timezone"`
		LocalTime     string  `json:"local_time"`
		LocalTimeUnix int64   `json:"local_time_unix"`
		IsInEU        bool    `json:"is_in_european_union"`
	} `json:"location"`
	Elapsed float64 `json:"elapsed_ms"`
}

// IPAPI implements the Provider interface for ipapi.is.
type IPAPI struct {
	provider.BaseProvider
}

// NewIPAPI creates a new IPAPI provider.
func NewIPAPI() *IPAPI {
	return &IPAPI{
		BaseProvider: provider.BaseProvider{
			ProviderName: "IP API",
			ProviderID:   "ipapi",
			URLTemplate:  "https://api.ipapi.is/?q={ip}",
		},
	}
}

// ParseResponse parses the ipapi.is API response into normalized fields.
func (i *IPAPI) ParseResponse(body []byte, statusCode int) (*provider.Result, error) {
	if statusCode != 200 {
		return provider.NewErrorResult(i, statusCode, fmt.Errorf("unexpected status code: %d", statusCode)), nil
	}

	var resp IPAPIResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	result := provider.NewSuccessResult(i, statusCode, resp)
	return result, nil
}

func init() {
	provider.Register(NewIPAPI())
}
