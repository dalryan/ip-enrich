package sources

// IPkey is a field from StopForumSpam.
//
// Fields:
//   - Value: The IP address as a string.
//   - Frequency: How often the IP address appears in spam databases.
//   - Appears: The number of different incidents involving this IP.
//   - ASN: The Autonomous System Number associated with the IP.
//   - Country: The country code from where the IP operates.
type IPkey struct {
	Value     string `json:"value"`
	Frequency int    `json:"frequency"`
	Appears   int    `json:"appears"`
	ASN       int    `json:"asn"`
	Country   string `json:"country"`
}

// StopForumSpamResponse is a response received from the StopForumSpam API.
// Fields:
//   - Success: An integer indicating the outcome of the API request.
//   - IP: The details of the IP address as defined in IPkey.
type StopForumSpamResponse struct {
	Success int   `json:"success"`
	IP      IPkey `json:"ip"`
}
