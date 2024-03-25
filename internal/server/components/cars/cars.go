package cars

import (
	"fmt"
	"strings"
)

var cars = map[string]Car{
	"id1": Init("Renault Logan"),
	"id2": Init("Renault Duster"),
	"id3": Init("BMW X6"),
	"id4": Init("BMW M5"),
	"id5": Init("VW Passat"),
	"id6": Init("VW Jetta"),
	"id7": Init("Audi A4"),
	"id8": Init("Audi Q7"),
}

type Car struct {
	Brand   string
	Name    string
	Counter int
}

func Init(name string) Car {
	splittedname := strings.Split(name, " ")
	brand, name := splittedname[0], splittedname[1]
	return Car{brand, name, 0}
}

func (c Car) String() string {
	return fmt.Sprintf("Brand: %v, Model: %v", c.Brand, c.Name)
}

func FindCarByID(id string) (Car, error) {
	if c, ok := cars[id]; ok {
		c.Counter++
		return c, nil
	}
	return Car{}, fmt.Errorf("not found")
}

func FindCarByBrand(brand string) []Car {
	var carModels []Car
	for _, model := range cars {
		if model.Brand == brand {
			carModels = append(carModels, model)
		}
	}
	return carModels
}

func FindCarByBrandAndModel(brand string, name string) []Car {
	var carModels []Car
	for _, model := range cars {
		if model.Brand == brand && model.Name == name {
			carModels = append(carModels, model)
		}
	}
	return carModels
}

func GetCarModels() []string {
	var carModels []string
	for _, model := range cars {
		carModels = append(carModels, model.String())
	}
	return carModels
}
