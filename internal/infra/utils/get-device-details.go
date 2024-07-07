package utils

import "github.com/mssola/useragent"

type DeviceDetails struct {
	Name     string
	Platform string
}

func GetDeviceDetails(userAgent string) DeviceDetails {
	ua := useragent.New(userAgent)

	name, _ := ua.Browser()
	device := DeviceDetails{
		Name:     name,
		Platform: ua.Platform(),
	}

	return device
}
