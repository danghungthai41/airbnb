package placemodel

import (
	"airbnb-golang/pkg/common"
	"errors"
	"strings"
)

type Place struct {
	common.SQLModel
	Owner_Id      int     `json:"ownerId" gorm:"column:owner_id"`
	CityId        int     `json:"cityId" gorm:"column:city_id"`
	Name          string  `json:"name" gorm:"column:name"`
	Address       string  `json:"address" gorm:"column:address"`
	Lat           float64 `json:"lat" gorm:"column:lat"`
	Lng           float64 `json:"lng" gorm:"column:lng"`
	PricePerNight float64 `json:"pricePerNight" gorm:"column:price_per_night"`
}

func (Place) TableName() string {
	return "places"
}
func (p *Place) Validate() error {
	p.Name = strings.TrimSpace(p.Name)
	p.Address = strings.TrimSpace(p.Address)

	if p.Name == "" || p.Address == "" {
		return errors.New("field can not be blank")
	}
	return nil
}
