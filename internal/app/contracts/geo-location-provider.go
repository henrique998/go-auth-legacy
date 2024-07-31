package contracts

type GeoLocationProvider interface {
	GetInfo(ip string) (string, string, error)
}
