package providers

import (
	"encoding/json"
	"fmt"

	"github.com/dalryan/ip-enrich/internal/provider"
)

// IPWhoisResponse represents the response from ipwho.is API.
type IPWhoisResponse struct {
	IP            string  `json:"ip"`
	Success       bool    `json:"success"`
	Type          string  `json:"type"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continent_code"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"region_code"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	IsEU          bool    `json:"is_eu"`
	Postal        string  `json:"postal"`
	CallingCode   string  `json:"calling_code"`
	Capital       string  `json:"capital"`
	Borders       string  `json:"borders"`
	Flag          struct {
		Img          string `json:"img"`
		Emoji        string `json:"emoji"`
		EmojiUnicode string `json:"emoji_unicode"`
	} `json:"flag"`
	Connection struct {
		ASN    int    `json:"asn"`
		Org    string `json:"org"`
		ISP    string `json:"isp"`
		Domain string `json:"domain"`
	} `json:"connection"`
	Timezone struct {
		ID          string `json:"id"`
		Abbr        string `json:"abbr"`
		IsDST       bool   `json:"is_dst"`
		Offset      int    `json:"offset"`
		UTC         string `json:"utc"`
		CurrentTime string `json:"current_time"`
	} `json:"timezone"`
}

// IPWhois implements the Provider interface for ipwho.is.
type IPWhois struct {
	provider.BaseProvider
}

// NewIPWhois creates a new IPWhois provider.
func NewIPWhois() *IPWhois {
	return &IPWhois{
		BaseProvider: provider.BaseProvider{
			ProviderName: "IP Whois",
			ProviderID:   "ipwhois",
			URLTemplate:  "https://ipwho.is/{ip}",
		},
	}
}

// ParseResponse parses the ipwho.is API response into normalized fields.
func (i *IPWhois) ParseResponse(body []byte, statusCode int) (*provider.Result, error) {
	if statusCode != 200 {
		return provider.NewErrorResult(i, statusCode, fmt.Errorf("unexpected status code: %d", statusCode)), nil
	}

	var resp IPWhoisResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !resp.Success {
		return provider.NewErrorResult(i, statusCode, fmt.Errorf("API returned success=false")), nil
	}

	result := provider.NewSuccessResult(i, statusCode, resp)

	return result, nil
}

func init() {
	provider.Register(NewIPWhois())
}
