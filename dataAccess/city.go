package dataAccess

type City struct {
	Name  string
	State string
	Years []AnnualCityData
}

func (c *City) FindDemographicsByYear(year int) *AnnualCityData {
	for _, d := range c.Years {
		if d.Year == year {
			return &d
		}
	}
	return nil
}

type AnnualCityData struct {
	Year       int
	Population int
}
