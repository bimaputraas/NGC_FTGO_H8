package main

import "fmt"

type Shape struct {
	ShapeType string
	Length    int
	Area      float64
}

func (s Shape) countarea(channel chan float64) {
	var output float64

	if s.ShapeType == "rectangle" {
		output = float64(s.Length * s.Length)
	} else if s.ShapeType == "circle" {
		output = float64(22 / 7 * s.Length / 2 * s.Length / 2)
	} else if s.ShapeType == "triangle" {
		output = float64((s.Length * s.Length) / 2)
	}

	channel <- output
}

func main() {
	channelrectangle := make(chan float64)
	defer close(channelrectangle)
	channelcircle := make(chan float64)
	defer close(channelcircle)
	channeltriangle := make(chan float64)
	defer close(channeltriangle)

	var areaResult float64

	input := []Shape{
		{ShapeType: "rectangle", Length: 5},
		{ShapeType: "circle", Length: 3},
		{ShapeType: "triangle", Length: 5},
		{ShapeType: "rectangle", Length: 15},
		{ShapeType: "circle", Length: 5},
	}

	for _, i := range input {
		if i.ShapeType == "rectangle" {
			go i.countarea(channelrectangle)
			areaResult = <-channelrectangle
		} else if i.ShapeType == "circle" {
			go i.countarea(channelcircle)
			areaResult = <-channelcircle
		} else if i.ShapeType == "triangle" {
			go i.countarea(channeltriangle)
			areaResult = <-channeltriangle
		}

		// fmt.Printf("Area from %s : %f\n", i.ShapeType, i.Area)
		fmt.Printf("Area from %s : %.2f\n", i.ShapeType, areaResult)
	}
	
	

}
