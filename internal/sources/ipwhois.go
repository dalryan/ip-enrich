package sources

type Flag struct {
	Img          string `json:"img"`
	Emoji        string `json:"emoji"`
	EmojiUnicode string `json:"emoji_unicode"`
}

type Connection struct {
	ASN    int    `json:"asn"`
	Org    string `json:"org"`
	ISP    string `json:"isp"`
	Domain string `json:"domain"`
}

type Timezone struct {
	ID          string `json:"id"`
	Abbr        string `json:"abbr"`
	IsDST       bool   `json:"is_dst"`
	Offset      int    `json:"offset"`
	UTC         string `json:"utc"`
	CurrentTime string `json:"current_time"`
}

// IPInfoResponse is a response received from the IPInfo API.
//
// Fields:
//   - IP: The IP address being queried.
//   - Success: Indicates whether the API request was successful.
//   - Type: The type of IP address (IPv4 or IPv6).
//   - Continent: The name of the continent.
//   - ContinentCode: The two-letter continent code.
//   - Country: The name of the country.
//   - CountryCode: The two-letter country code.
//   - Region: The name of the region.
//   - RegionCode: The two-letter region code.
//   - City: The name of the city.
//   - Latitude: The latitude of the location.
//   - Longitude: The longitude of the location.
//   - IsEU: Indicates whether the IP address is in the European Union.
//   - Postal: The postal code of the location.
//   - CallingCode: The international calling code for the country.
//   - Capital: The capital city of the country.
//   - Borders: A list of bordering countries.
//   - Flag: A Flag struct containing image and emoji details.
//   - Connection: A Connection struct containing ASN, organization, ISP, and domain details.
//   - Timezone: A Timezone struct containing timezone details.
type IPInfoResponse struct {
	IP            string     `json:"ip"`
	Success       bool       `json:"success"`
	Type          string     `json:"type"`
	Continent     string     `json:"continent"`
	ContinentCode string     `json:"continent_code"`
	Country       string     `json:"country"`
	CountryCode   string     `json:"country_code"`
	Region        string     `json:"region"`
	RegionCode    string     `json:"region_code"`
	City          string     `json:"city"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	IsEU          bool       `json:"is_eu"`
	Postal        string     `json:"postal"`
	CallingCode   string     `json:"calling_code"`
	Capital       string     `json:"capital"`
	Borders       string     `json:"borders"`
	Flag          Flag       `json:"flag"`
	Connection    Connection `json:"connection"`
	Timezone      Timezone   `json:"timezone"`
}
