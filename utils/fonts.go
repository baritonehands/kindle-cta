package utils

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func LoadFont(name string) (*truetype.Font, error) {
	fontBytes, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	return f, nil
}

const spacing = 1.5

func WriteTextToFile(font *truetype.Font, size float64, text string) error {
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))

	err := DrawTextToImage(rgba, font, size, text)
	if err != nil {
		return err
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		return err
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	fmt.Println("Wrote out.png OK.")
	return nil
}

func DrawTextToImage(rgba draw.Image, font *truetype.Font, size float64, text string) error {
	// Initialize the context.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}

	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(167)
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(size)>>6))
	_, err := c.DrawString(text, pt)
	if err != nil {
		return err
	}
	pt.Y += c.PointToFixed(size * spacing)
	return nil
}
