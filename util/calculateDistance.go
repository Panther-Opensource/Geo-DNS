package util

import (
	"Geo-DNS/stores"
	"math"
)

func CalculateDistance(ip string, ip2 string) float64 {
	ip1loc, e := stores.Db.Get_all(ip)
	if e != nil {
		return -1
	}
	ip2loc, e := stores.Db.Get_all(ip2)
	if e != nil {
		return -1
	}
	distance := getDistanceFromLatLonInKm(float64(ip1loc.Latitude), float64(ip1loc.Longitude), float64(ip2loc.Latitude), float64(ip2loc.Longitude))
	return distance
}
func getDistanceFromLatLonInKm(lat1, lon1, lat2, lon2 float64) float64 {
	R := 6371.0
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := R * c
	return d
}
func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
