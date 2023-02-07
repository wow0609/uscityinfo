package dataAccess

type CityRepository interface {
	CityByName(cityName string) *City
	CityByNameYear(cityName string, year int) *AnnualCityData
}
