package main

import "log"

func main() {
	triangle := TriangularObject{
		height: 1,
		base:   2,
		length: 3,
	}

	rectangle := RectangleObject{
		x: 23,
		y: 4,
		z: 1,
	}
	x := 1
	var object Object
	if x > 2 {
		object = triangle
	} else {
		object = rectangle
	}

	volumeCalculator := VolumeCalculator{
		object: object,
	}

	volume := volumeCalculator.Volume()
	log.Print(volume)

}

type Object interface {
	Perimeter() float64
	Height() float64
}

type VolumeCalculator struct {
	object Object
}

func (calculator VolumeCalculator) Volume() float64 {
	return calculator.object.Perimeter() * calculator.object.Height()
}

type TriangularObject struct {
	height float64
	base   float64
	length float64
}

func (t TriangularObject) Perimeter() float64 {
	return t.height * t.base / 2
}

func (t TriangularObject) Height() float64 {
	return t.length
}

type RectangleObject struct {
	x float64
	y float64
	z float64
}

func (r RectangleObject) Perimeter() float64 {
	return r.x * r.y
}

func (r RectangleObject) Height() float64 {
	return r.z
}
