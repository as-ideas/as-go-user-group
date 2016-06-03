package main

import (
	"image/png"
	"os"
	"fmt"
	"image"
)

func main() {

	filename := "/Users/swaschni/Projekte/GO-IDEAS/mate/vinyl.png"
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	// Must specifically use jpeg.Decode() or it
	// would encounter unknown format error
	img, err := png.Decode(file)
	check(err)

	lastPoint := image.Point{0, 0}
	nextPoint := image.Point{251, 128}
	var amplitudes []uint8

	for nextPoint != nil {
		at := img.At(nextPoint.X, nextPoint.Y)
		value, _, _, _ := at.RGBA()

		amplitudes = append(amplitudes, value)

		// TODO -> Save value
		fmt.Printf("Stuff %b", value)

		// find next point
		allNextPoints := getNeighbours(nextPoint)
		fmt.Printf("%v", allNextPoints)

		nextPoint = getNextPoint(img, allNextPoints, lastPoint)

		lastPoint = nextPoint
	}

	// TODO Play Sound
}

func getNextPoint(image image.Image, neighbours []image.Point, lastPoint image.Point) image.Point {
	var nextPoint image.Point
	for _, element := range neighbours {
		if (element != lastPoint && image.At(element.X, element.Y).RGBA()[0] != 0) {
			nextPoint = element
		}
	}
	return nextPoint
}

func getNeighbours(point image.Point) []image.Point {
	var res []image.Point

	res = append(res, point.Add(image.Point{-1,-1}))
	res = append(res, point.Add(image.Point{0,-1}))
	res = append(res, point.Add(image.Point{1,-1}))
	res = append(res, point.Add(image.Point{1,0}))
	res = append(res, point.Add(image.Point{1,1}))
	res = append(res, point.Add(image.Point{0,1}))
	res = append(res, point.Add(image.Point{-1,-1}))
	res = append(res, point.Add(image.Point{-1,0}))

	return res
}

func check(myError error) {
	if myError != nil {
		panic(myError)
	}
}