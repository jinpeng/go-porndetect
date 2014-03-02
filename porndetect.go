package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	var filepath = flag.String("input", "", "Image to be checked.")
	flag.Parse()

	im, _, err := decode(*filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	isPorn := imageCheck(im)
	if isPorn {
		fmt.Println(*filepath + "is a porn image.")
	} else {
		fmt.Println(*filepath + "is NOT a porn image.")
	}
}

func decode(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(bufio.NewReader(f))
}

func colorCheck(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	_, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	return (86 <= cb) && (cb <= 117) && (140 <= cr) && (cr <= 168)
}

func imageCheck(im image.Image) bool {
	count := 0
	b := im.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if colorCheck(im.At(x, y)) {
				count++
			}
		}
	}
	return float64(count) > float64(b.Max.X)*float64(b.Max.Y)*float64(0.3)
}
