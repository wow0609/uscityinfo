package dataAccess

type CityRepository interface {
	AllCities() []*City
	CityByName(cityName string) *City
	CityByNameYear(cityName string, year int) *AnnualCityData
}
