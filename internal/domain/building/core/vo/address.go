package vo

import "fmt"

type Address struct {
	Country  string      `json:"country"`
	Region   string      `json:"region"`
	Province AddressPair `json:"province"`
	District AddressPair `json:"district"`
}

type AddressPair struct {
	Id    uint   `json:"id"`
	Value string `json:"value"`
}

func NewAddress(country, region string, province, district AddressPair) *Address {
	return &Address{
		Country:  country,
		Region:   region,
		Province: province,
		District: district,
	}
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Country, a.Region, a.Province.Value, a.District.Value)
}
