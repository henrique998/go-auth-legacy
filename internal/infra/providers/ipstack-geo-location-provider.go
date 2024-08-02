package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPStackGeoLocationProvider struct {
	APiKey string
}

type ipstackResponse struct {
	CountryName string `json:"country_name"`
	City        string `json:"city"`
}

func (p *IPStackGeoLocationProvider) GetInfo(ip string) (string, string, error) {
	url := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip, p.APiKey)

	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to get geolocation data: %s", res.Status)
	}

	var result ipstackResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", "", err
	}

	return result.CountryName, result.City, nil
}
