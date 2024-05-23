package sources

// GreyNoiseCommunityResponse is a response from the GreyNoise Community API.
//
// Fields:
//   - IP: The IP address being queried.
//   - Noise: Indicates whether the IP is labeled as "noise" (background traffic).
//   - Riot: Indicates whether the IP is part of a "riot" (benign data centers, CDNs, etc.).
//   - Classification: A string that categorizes the type of noise.
//   - Name: A descriptive name for the classification.
//   - Link: A URL to GreyNoise for detailed information about the IP address.
//   - LastSeen: The last date and time the IP was observed by GreyNoise.
//   - Message: A message providing additional context or information about the IP.
type GreyNoiseCommunityResponse struct {
	IP             string `json:"ip"`
	Noise          bool   `json:"noise"`
	Riot           bool   `json:"riot"`
	Classification string `json:"classification"`
	Name           string `json:"name"`
	Link           string `json:"link"`
	LastSeen       string `json:"last_seen"`
	Message        string `json:"message"`
}
