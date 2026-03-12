package ui

import (
	"fmt"
	"image/color"

	"github.com/baritonehands/kindle-cta/utils"
	"github.com/golang/freetype/truetype"
	"github.com/simsor/go-kindle/kindle"
)

var FontFreeSansRegular *truetype.Font = utils.LoadFont("assets/FreeSans.ttf")
var FontFreeSansBold *truetype.Font = utils.LoadFont("assets/FreeSansBold.ttf")

var Regular8PtBlack = NewFontRenderer(FontFreeSansRegular, kindle.Framebuffer(), 8)
var Regular12PtBlack = NewFontRenderer(FontFreeSansRegular, kindle.Framebuffer(), 12)
var Regular12PtWhite = NewFontRenderer(FontFreeSansRegular, kindle.Framebuffer(), 12)
var Bold12PtBlack = NewFontRenderer(FontFreeSansBold, kindle.Framebuffer(), 12)
var Bold16PtBlack = NewFontRenderer(FontFreeSansBold, kindle.Framebuffer(), 16)

func init() {
	fmt.Println("Initializing constants")
	Regular12PtWhite.SetFontColor(color.White)
}
