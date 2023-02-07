package dataAccess

type City struct {
	Name  string
	State string
	Years []AnnualCityData
}

type AnnualCityData struct {
	Year       int
	Population int
}
