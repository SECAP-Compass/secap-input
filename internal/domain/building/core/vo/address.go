package vo

import "fmt"

type Address struct {
	Country  string `json:"country"`
	Region   string `json:"region"`
	Province string `json:"province"`
	District string `json:"district"`
}

func NewAddress(country, region, province, district, postalCode string) *Address {
	return &Address{
		Country:  country,
		Region:   region,
		Province: province,
		District: district,
	}
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Country, a.Region, a.Province, a.District)
}
