package dataAccess

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type DataApiCityRepository struct {
	CityRepository
	queryAll string
	queryUrl string
}

func NewDataApiCityData() *DataApiCityRepository {
	return &DataApiCityRepository{
		queryAll: "https://public.opendatasoft.com/api/records/1.0/search/?dataset=us-cities-demographics&rows=1000",
		queryUrl: "https://public.opendatasoft.com/api/records/1.0/search/?dataset=us-cities-demographics&q=city%3D{{cityName}}&rows=1",
	}
}

type DataApiStateListDto struct {
	Records []struct {
		DatasetId string `json:"datasetid"`
		RecordId  string `json:"recordid"`
		Fields    struct {
			State          string `json:"state_code"`
			City           string `json:"city"`
			Population     int    `json:"total_population"`
			MalePopulation int    `json:"male_population"`
		} `json:"fields"`
	} `json:"records"`
}

func (l *DataApiStateListDto) allCities() (cities []*City) {
	for _, v := range l.Records {
		city := &City{}
		city.Name = v.Fields.City
		city.State = v.Fields.State
		city.Years = append(city.Years, AnnualCityData{
			Year:       2023,
			Population: v.Fields.Population,
		})
		cities = append(cities, city)
	}
	return cities

}
func (l *DataApiStateListDto) findCity(cityName string) *City {
	city := &City{}
	for _, v := range l.Records {
		if v.Fields.City == cityName {
			city.Name = v.Fields.City
			city.State = v.Fields.State
			city.Years = append(city.Years, AnnualCityData{
				Year:       2023,
				Population: v.Fields.Population,
			})
			return city
		}
	}
	return nil
}

// type DataApiStateDto struct {
// 	idState    string `json: "ID State"`
// 	state      string `json: "State"`
// 	idYear     int    `json: "ID Year"`
// 	population int    `json: "Population"`
// 	slugState  string `json: "Slug State"`
// }
func getJsonFromEndpoint(url string, target interface{}) error {
	client := &http.Client{
		Timeout: 12 * time.Second,
	}

	res, err := client.Get(url)

	if err != nil {
		return err
	}
	//result := &DataApiStateListDto{}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataApiCityRepository) CityByName(cityName string) *City {
	getUrl := strings.Replace(d.queryUrl, "{{cityName}}", cityName, 1)

	result := &DataApiStateListDto{}
	getJsonFromEndpoint(getUrl, result)

	city := result.findCity(cityName)
	if city == nil {
		return nil
	}

	log.Printf("City: [%s]", city.Name)
	return city

}

func (d *DataApiCityRepository) AllCities() []*City {
	result := &DataApiStateListDto{}
	getJsonFromEndpoint(d.queryAll, result)
	//log.Print(result)
	return result.allCities()

}

func (d *DataApiCityRepository) CityByNameYear(cityName string, year int) *AnnualCityData {
	if city := d.CityByName(cityName); city == nil {
		return nil
	} else {
		return &city.Years[0]
	}

}
