package main

import "fmt"

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
		if t == Inch && u.T == CM {
			value = value / 2.54
		} else if t == CM && u.T == Inch {
			value = value * 2.54
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type DimensionsInInches struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsInInches) Length() Unit {
	return d.length
}

func (d DimensionsInInches) Width() Unit {
	return d.width
}

func (d DimensionsInInches) Height() Unit {
	return d.height
}

type DimensionsInCM struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsInCM) Length() Unit {
	return d.length
}

func (d DimensionsInCM) Width() Unit {
	return d.width
}

func (d DimensionsInCM) Height() Unit {
	return d.height
}

type BMW struct{}

func (b BMW) Brand() string {
	return "BMW"
}

func (b BMW) Model() string {
	return "X5"
}

func (b BMW) Dimensions() Dimensions {
	return DimensionsInCM{
		length: Unit{Value: 492.2, T: CM},
		width:  Unit{Value: 200.4, T: CM},
		height: Unit{Value: 174.5, T: CM},
	}
}

func (b BMW) MaxSpeed() int {
	return 250
}

func (b BMW) EnginePower() int {
	return 340
}

type Mercedes struct{}

func (m Mercedes) Brand() string {
	return "Mercedes"
}

func (m Mercedes) Model() string {
	return "GLE"
}

func (m Mercedes) Dimensions() Dimensions {
	return DimensionsInCM{
		length: Unit{Value: 481.9, T: CM},
		width:  Unit{Value: 193.5, T: CM},
		height: Unit{Value: 179.6, T: CM},
	}
}

func (m Mercedes) MaxSpeed() int {
	return 240
}

func (m Mercedes) EnginePower() int {
	return 367
}

type Dodge struct{}

func (d Dodge) Brand() string {
	return "Dodge"
}

func (d Dodge) Model() string {
	return "Charger"
}

func (d Dodge) Dimensions() Dimensions {
	return DimensionsInInches{
		length: Unit{Value: 198.4, T: Inch},
		width:  Unit{Value: 75.0, T: Inch},
		height: Unit{Value: 57.8, T: Inch},
	}
}

func (d Dodge) MaxSpeed() int {
	return 203
}

func (d Dodge) EnginePower() int {
	return 292
}

func main() {
	bmw := BMW{}
	mercedes := Mercedes{}
	dodge := Dodge{}

	autos := []Auto{bmw, mercedes, dodge}

	for _, auto := range autos {
		fmt.Printf("Brand: %s, Model: %s\n", auto.Brand(), auto.Model())
		fmt.Printf("Dimensions (LxWxH): %.2f x %.2f x %.2f %s\n",
			auto.Dimensions().Length().Value,
			auto.Dimensions().Width().Value,
			auto.Dimensions().Height().Value,
			auto.Dimensions().Length().T)
		fmt.Printf("Max Speed: %d km/h, Engine Power: %d HP\n\n",
			auto.MaxSpeed(), auto.EnginePower())
	}
}
