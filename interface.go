package figuras 

import "fmt"

type Geometria interface {
	area() float64
	perimetro() float64
}

func Calcular(geo Geometria) {
	fmt.Println("Area: ", geo.area())
	fmt.Println("Perimetro: ", geo.perimetro())
}