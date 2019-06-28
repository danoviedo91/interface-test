package main

import (
	"fmt"

	"github.com/danoviedo91/test/models"
)

// Shape ...
type Shape interface {
	ShapeType() string
	Area() float64
	Perimeter() float64
}

func main() {

	shapes := []Shape{}
	shapes = append(shapes, models.Square{Side: 4.3})
	shapes = append(shapes, models.Rectangle{Width: 5, Length: 4})

	for _, shape := range shapes {
		fmt.Printf("%v area: %v cm2\n", shape.ShapeType(), shape.Area())
		fmt.Printf("%v perimeter: %v cm\n", shape.ShapeType(), shape.Perimeter())
		fmt.Println("---------------------")
	}

}
