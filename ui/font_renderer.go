package ui

import (
	"image"
	"image/draw"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const dpi = 167 // Kindle 4

type FontRenderer struct {
	context       *freetype.Context
	Size, Spacing float64
}

func NewFontRenderer(font *truetype.Font, dst draw.Image, size float64) *FontRenderer {
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetSrc(image.Black)
	return &FontRenderer{context: c, Size: size, Spacing: 1.5}
}

func (r *FontRenderer) PrintAt(x, y int, text string) error {
	lines := strings.Split(text, "\n")
	// Draw the text.
	pt := freetype.Pt(x, y+int(r.context.PointToFixed(r.Size)>>6))
	for _, s := range lines {
		_, err := r.context.DrawString(s, pt)
		if err != nil {
			return err
		}
		pt.Y += r.context.PointToFixed(r.Size * r.Spacing)
	}
	return nil
}

func (r *FontRenderer) CharHeight() int {
	return r.context.PointToFixed(r.Size * r.Spacing).Round()
}
