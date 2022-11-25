package main

import "fmt"

/*
	Implement Adapter pattern
*/

// My example:
// Given Pixel structure that contain some color in BGR format
// But we want RGB format, so we need to create adapter for Pixel

// given color format
type BGR struct {
	Blue  int
	Green int
	Red   int
}

// given pixel struct
type Pixel struct {
	x, y  int
	color BGR
}

// Pixel struct color getter
func (p *Pixel) Color() BGR {
	return p.color
}

// desired color format
type RGB struct {
	Red   int
	Green int
	Blue  int
}

// Pixel adapter struct
type PixelAdapter struct {
	pixel *Pixel
}

func (pa *PixelAdapter) Color() RGB {
	return RGB{
		Red:   pa.pixel.color.Red,
		Green: pa.pixel.color.Green,
		Blue:  pa.pixel.color.Blue,
	}
}

func main() {
	// given pixel with BGR color
	pixel := Pixel{
		x: 10,
		y: 10,
		color: BGR{
			Blue:  255,
			Green: 0,
			Red:   0,
		},
	}

	// create adapter for pixel
	pixelAdapter := PixelAdapter{pixel: &pixel}

	// print original pixel color
	fmt.Println(pixel.Color())

	// print pixel color using adapter
	fmt.Println(pixelAdapter.Color())
}
