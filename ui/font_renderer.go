package ui

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const dpi = 167 // Kindle 4

type FontRenderer struct {
	context       *freetype.Context
	size, spacing float64
}

func NewFontRenderer(font *truetype.Font, size float64) *FontRenderer {
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetSrc(image.Black)
	return &FontRenderer{context: c, size: size, spacing: 1.5}
}

func (r *FontRenderer) PrintAt(dst draw.Image, x, y int, text string) error {
	r.context.SetDst(dst)
	r.context.SetClip(dst.Bounds())
	lines := strings.Split(text, "\n")
	// Draw the text.
	pt := freetype.Pt(x, y+int(r.context.PointToFixed(r.size)>>6))
	for _, s := range lines {
		if Debug {
			fmt.Printf("Drawing text at: %v\n", pt)
		}
		_, err := r.context.DrawString(s, pt)
		if err != nil {
			return err
		}
		pt.Y += r.context.PointToFixed(r.size * r.spacing)
	}
	return nil
}

func (r *FontRenderer) CharHeight() int {
	return r.context.PointToFixed(r.size * r.spacing).Round()
}

func (r *FontRenderer) SetFontSize(fontSize float64) {
	r.size = fontSize
	r.context.SetFontSize(fontSize)
}

func (r *FontRenderer) SetFontColor(color color.Color) {
	r.context.SetSrc(image.NewUniform(color))
}
