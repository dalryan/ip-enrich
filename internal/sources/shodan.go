package sources

// ShodanResponse is a response received from the Shodan API.
//
// Fields:
//   - CPEs: A list of Common Platform Enumeration (CPE) strings.
//   - Hostnames: A list of hostnames associated with the IP address.
//   - IP: The IP address being queried.
//   - Ports: A list of open ports on the IP address.
//   - Tags: A list of tags associated with the IP address.
//   - Vulns: A list of vulnerabilities associated with the IP address.
type ShodanResponse struct {
	CPEs      []string `json:"cpes"`
	Hostnames []string `json:"hostnames"`
	IP        string   `json:"ip"`
	Ports     []int    `json:"ports"`
	Tags      []string `json:"tags"`
	Vulns     []string `json:"vulns"`
}
