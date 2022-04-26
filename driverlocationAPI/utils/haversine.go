package utils

import (
	"math"
)

func Haversine(latitude1 float64, longitude1 float64, latitude2 float64, longitude2 float64) float64 {
	var rad, c float64
	dLat := (math.Pi / 180) * (latitude2 - latitude1)
	dLon := (math.Pi / 180) * (longitude2 - longitude1)

	latitude1 = (math.Pi / 180) * latitude1
	latitude2 = (math.Pi / 180) * latitude2

	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Pow(math.Sin(dLon/2), 2)*
			math.Cos(latitude1)*math.Cos(latitude2)
	rad = 6371
	c = 2 * math.Asin(math.Sqrt(a))
	return math.Round((rad*c)*100) / 100
}
