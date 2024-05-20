package models

type IpHopsReqBody struct {
	Hostname   string `json:"hostname"`
	NumQueries int    `json:"numQueries"`
	WaitTime   int    `json:"waitTime"`
	MaxHops    int    `json:"maxHops"`
}

type IpHop struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
	Rtt      string `json:"returnTime"`
}
