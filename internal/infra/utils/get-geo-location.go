package utils

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

func GetGeoLocation(ip string) (string, string, error) {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	ipAddress := net.ParseIP(ip)
	record, err := db.City(ipAddress)
	if err != nil {
		return "", "", err
	}

	country := record.Country.Names["en"]
	city := record.City.Names["en"]

	return country, city, nil
}
