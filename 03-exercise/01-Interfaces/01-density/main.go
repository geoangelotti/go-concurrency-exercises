package main

import "fmt"

// Metal - mass and volume information
type Metal struct {
	mass   float64
	volume float64
}

// Gas - measurements for Gas
type Gas struct {
	pressure      float64
	temperature   float64
	molecularMass float64
}

// Density - return density of metal
func (m *Metal) Density() float64 {
	// density = mass/volume
	return m.mass / m.volume
}

// Density - return density of liquid
func (g *Gas) Density() float64 {
	return (g.molecularMass * g.pressure) / (0.0821 * (g.temperature + 273.0))
}

type Dense interface {
	Density() float64
}

// IsDenser - compare density of two objects
func IsDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {
	gold := Metal{478, 24}
	silver := Metal{100, 10}

	result := IsDenser(&gold, &silver)
	if result {
		fmt.Println("gold has higher density than silver")
	} else {
		fmt.Println("silver has higher density than gold")
	}

	oxygen := Gas{pressure: 5, temperature: 27, molecularMass: 32}
	hydrogen := Gas{pressure: 1, temperature: 0, molecularMass: 2}

	result = IsDenser(&oxygen, &hydrogen)
	if result {
		fmt.Println("oxygen has higher density than hydrogen")
	} else {
		fmt.Println("hydrogen has higher density than oxygen")
	}
}
