package dataAccess

type DummyCityRepository struct {
	CityRepository
}

func NewDummyCityData() *DummyCityRepository {
	return &DummyCityRepository{}
	//data.cityData = append(data.cityData, *Birmingham())
	//data.cityData = append(data.cityData, *Mobile())
	//return data
}

func (d *DummyCityRepository) CityByName(cityName string) *City {

	switch cityName {
	case "Birmingham":
		return Birmingham()
	case "Mobile":
		return Mobile()
	default:
		return nil
	}
}

func (d *DummyCityRepository) AllCities() []*City {
	return []*City{Birmingham(), Mobile()}
}

func (d *DummyCityRepository) CityByNameYear(cityName string, year int) *AnnualCityData {
	if city := d.CityByName(cityName); city == nil {
		return nil
	} else {
		return &city.Years[0]
	}

}

//Database
func Birmingham() *City {
	return &City{
		Name:  "Birmingham",
		State: "AL",
		Years: []AnnualCityData{
			{
				Year:       1980,
				Population: 111_980,
			},
			{
				Year:       2000,
				Population: 222_000,
			},
			{
				Year:       2020,
				Population: 300_000,
			},
		},
	}
}

func Mobile() *City {
	return &City{
		Name:  "Mobile",
		State: "AL",
		Years: []AnnualCityData{
			{
				Year:       1980,
				Population: 1_980,
			},
			{
				Year:       2000,
				Population: 2_000,
			},
			{
				Year:       2020,
				Population: 3_000,
			},
		},
	}
}
