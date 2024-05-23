package sources

type Company struct {
	Name        string `json:"name"`
	AbuserScore string `json:"abuser_score"`
	Domain      string `json:"domain"`
	Type        string `json:"type"`
	Network     string `json:"network"`
	Whois       string `json:"whois"`
}

type Datacenter struct {
	Name    string `json:"datacenter"`
	Domain  string `json:"domain"`
	Network string `json:"network"`
}

type ASN struct {
	ASN         int    `json:"asn"`
	AbuserScore string `json:"abuser_score"`
	Route       string `json:"route"`
	Descr       string `json:"descr"`
	Country     string `json:"country"`
	Active      bool   `json:"active"`
	Org         string `json:"org"`
	Domain      string `json:"domain"`
	Abuse       string `json:"abuse"`
	Type        string `json:"type"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	RIR         string `json:"rir"`
	Whois       string `json:"whois"`
}

type Location struct {
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
	IsDST         bool    `json:"is_dst"`
}

// IPAPIResponse is a response received from the IPAPI API.
//
// Fields:
//   - IP: The IP address being queried.
//   - RIR: The Regional Internet Registry associated with the IP.
//   - IsBogon: Indicates whether the IP is a bogon (unallocated or reserved address space).
//   - IsMobile: Indicates whether the IP is associated with a mobile network.
//   - IsCrawler: Indicates whether the IP is a web crawler.
//   - IsDatacenter: Indicates whether the IP is part of a datacenter.
//   - IsTor: Indicates whether the IP is part of the Tor network.
//   - IsProxy: Indicates whether the IP is a proxy server.
//   - IsVPN: Indicates whether the IP is a VPN server.
//   - IsAbuser: Indicates whether the IP is associated with abusive behavior.
//   - Company: A Company struct containing details about the IP's company.
//   - Datacenter: A Datacenter struct containing details about the IP's datacenter.
//   - ASN: An ASN struct containing details about the IP's Autonomous System Number.
//   - Location: A Location struct containing details about the IP's location.
type IPAPIResponse struct {
	IP           string     `json:"ip"`
	RIR          string     `json:"rir"`
	IsBogon      bool       `json:"is_bogon"`
	IsMobile     bool       `json:"is_mobile"`
	IsCrawler    bool       `json:"is_crawler"`
	IsDatacenter bool       `json:"is_datacenter"`
	IsTor        bool       `json:"is_tor"`
	IsProxy      bool       `json:"is_proxy"`
	IsVPN        bool       `json:"is_vpn"`
	IsAbuser     bool       `json:"is_abuser"`
	Company      Company    `json:"company"`
	Datacenter   Datacenter `json:"datacenter"`
	ASN          ASN        `json:"asn"`
	Location     Location   `json:"location"`
}
