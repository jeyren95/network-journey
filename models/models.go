package models

type IpHopsReqBody struct {
	Hostname   string `json:"hostname"`
	WaitTime   int    `json:"waitTime"`
	MaxHops    int    `json:"maxHops"`
}

type IpHop struct {
	Hostname    string   `json:"hostname"`
	Ip          string   `json:"ip"`
	IsIpPrivate bool			`json:"isIpPrivate"`
	ReturnTime string `json:"returnTime"`
}

type Geolocation struct {
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}
