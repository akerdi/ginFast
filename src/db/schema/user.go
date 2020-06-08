package schema

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type User struct {
	Name string `gorm:"type:varchar(120);not null;" json:"name"`
	ID int `gorm:"primary_key" json:"id"`
	Houses Houses `gorm:"type:LONGTEXT;not null"`
}

type House struct {
	Address string `json:"address"`
	Price int `json:"price"`
}

type Houses []House

func (aHouse House) Value() (driver.Value, error)  {
	bytes, err := json.Marshal(aHouse)
	return string(bytes), err
}
func (aHouse House) Scan(src interface{}) error {
	switch value := src.(type) {
	case	string:
		return json.Unmarshal([]byte(value), aHouse)
	case []byte:
		return json.Unmarshal(value, aHouse)
	default:
		return errors.New("not supported")
	}
}