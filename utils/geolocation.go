package utils

import (
	"fmt"
	"net"
	"net/http"
	"io"
	"encoding/json"
	"github.com/jeyren95/network-journey/models"
)

const API_ENDPOINT = "http://ip-api.com/json"

func GetGeolocations(ipHops []models.IpHop) ([]models.Geolocation, error) {
	var geolocations []models.Geolocation
	for _, hop := range(ipHops) {
		if hop.IsIpPrivate {
			continue
		}

		endpoint := fmt.Sprint(API_ENDPOINT, "/", hop.Ip)
		resp, err := http.Get(endpoint)

		if err != nil {
			return nil, err
		}

		// resp.Body implements reader interface
		// use io to read the data into a byte slice
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			return nil, err
		}

		// unmarshal json encoded data
		var data models.Geolocation
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		geolocations = append(geolocations, data)
	}
	
	return geolocations, nil
}

func isPrivate(ip string) bool {
	parsedIp := net.ParseIP(ip)
	return parsedIp.IsPrivate()
}

