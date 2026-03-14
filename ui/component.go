package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/simsor/go-kindle/framebuffer"
)

type Component struct {
	Bounds          image.Rectangle
	Padding, Margin int
	BorderColor     color.Color
	visible, dirty  bool
}

func NewComponent(x, y int, width, height int) Component {
	pos := image.Pt(x, y)
	return Component{
		Bounds: image.Rectangle{
			Min: pos,
			Max: image.Pt(x+width, y+height),
		},
		BorderColor: color.Black,
		visible:     true,
		dirty:       true,
	}
}

func (c *Component) Translate(pos image.Point) image.Point {
	return c.Bounds.Min.Add(pos).Add(image.Pt(c.Margin+c.Padding, c.Margin+c.Padding))
}

func (c *Component) clear(device *framebuffer.Device) {
	for y := c.Bounds.Min.Y; y < c.Bounds.Max.Y; y++ {
		for x := c.Bounds.Min.X; x < c.Bounds.Max.X; x++ {
			device.Set(x, y, color.White)
		}
	}
}

func (c *Component) Render(device *framebuffer.Device) {
	if !c.visible {
		if c.dirty {
			c.clear(device)
			c.dirty = false
		}
		return
	}

	if !c.dirty {
		return
	}

	xMin := c.Bounds.Min.X + c.Margin
	xMax := c.Bounds.Max.X - c.Margin
	yMin := c.Bounds.Min.Y + c.Margin
	yMax := c.Bounds.Max.Y - c.Margin

	fmt.Printf("Printing vertical border at %d,%d\n", xMin, xMax)
	for y := yMin; y < yMax; y++ {
		if y == yMin || y == yMax-1 {
			fmt.Printf("Printing horizontal border at %d\n", y)
			for x := xMin; x < xMax; x++ {
				device.Set(x, y, c.BorderColor)
			}
		} else {
			device.Set(xMin, y, c.BorderColor)
			device.Set(xMax-1, y, c.BorderColor)
		}
	}
}

func (c *Component) hide() {
	prev := c.visible
	c.visible = false
	c.dirty = c.dirty || prev != c.visible
}

func (c *Component) show() {
	prev := c.visible
	c.visible = true
	c.dirty = c.dirty || prev != c.visible
}

type Renderable interface {
	Render(device *framebuffer.Device)
}
