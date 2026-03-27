package ui

import (
	"fmt"
	"image/color"

	"github.com/baritonehands/kindle-cta/utils"
)

var FontFreeSansRegular = utils.LoadFont("assets/FreeSans.ttf")
var FontFreeSansBold = utils.LoadFont("assets/FreeSansBold.ttf")

var Regular8PtBlack = NewFontRenderer(FontFreeSansRegular, 8)
var Regular12PtBlack = NewFontRenderer(FontFreeSansRegular, 12)
var Regular12PtWhite = NewFontRenderer(FontFreeSansRegular, 12)
var Bold12PtBlack = NewFontRenderer(FontFreeSansBold, 12)
var Bold16PtBlack = NewFontRenderer(FontFreeSansBold, 16)

const Debug = false

func init() {
	fmt.Println("Initializing constants")
	Regular12PtWhite.SetFontColor(color.White)
}
