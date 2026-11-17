// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package render

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging/interpolation"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
	"github.com/cocosip/go-dicom/pkg/imaging/transform"
)

// Graphic draws an overlay in final viewport coordinates.
type Graphic interface {
	Bounds() image.Rectangle
	Draw(draw.Image) error
}

// ImageGraphic draws Image with its upper-left corner at Position.
type ImageGraphic struct {
	Image    image.Image
	Position image.Point
}

// Bounds returns the destination rectangle occupied by ImageGraphic.
func (g ImageGraphic) Bounds() image.Rectangle {
	if g.Image == nil {
		return image.Rectangle{}
	}
	return g.Image.Bounds().Add(g.Position.Sub(g.Image.Bounds().Min))
}

// Draw composites ImageGraphic over dst.
func (g ImageGraphic) Draw(dst draw.Image) error {
	if dst == nil || g.Image == nil {
		return nil
	}
	draw.Draw(dst, g.Bounds(), g.Image, g.Image.Bounds().Min, draw.Over)
	return nil
}

// CompositeGraphic draws graphics in slice order.
type CompositeGraphic struct{ Graphics []Graphic }

// Bounds returns the union of all non-nil child graphic bounds.
func (g CompositeGraphic) Bounds() image.Rectangle {
	var bounds image.Rectangle
	for _, graphic := range g.Graphics {
		if graphic == nil {
			continue
		}
		if bounds.Empty() {
			bounds = graphic.Bounds()
		} else {
			bounds = bounds.Union(graphic.Bounds())
		}
	}
	return bounds
}

// Draw composites each graphic over dst in slice order.
func (g CompositeGraphic) Draw(dst draw.Image) error {
	for _, graphic := range g.Graphics {
		if graphic == nil {
			continue
		}
		if err := graphic.Draw(dst); err != nil {
			return err
		}
	}
	return nil
}

// DrawGraphics copies source and draws graphics in stable order.
func DrawGraphics(source image.Image, graphics []Graphic) (*image.RGBA, error) {
	if source == nil {
		return nil, fmt.Errorf("source image is nil")
	}
	destination := image.NewRGBA(source.Bounds())
	draw.Draw(destination, destination.Bounds(), source, source.Bounds().Min, draw.Src)
	if err := (CompositeGraphic{Graphics: graphics}).Draw(destination); err != nil {
		return nil, err
	}
	return destination, nil
}

// ApplyAffine renders source through a destination-to-source affine map. A
// zero viewport uses the normalized transformed bounds; a non-zero viewport is
// used verbatim and samples outside source with background.
func ApplyAffine(source image.Image, matrix transform.Affine2D, viewport image.Rectangle, mode interpolation.Mode, background color.Color) (*image.RGBA, error) {
	if source == nil {
		return nil, fmt.Errorf("source image is nil")
	}
	if mode != interpolation.ModeNearestNeighbor && mode != interpolation.ModeBilinear {
		return nil, fmt.Errorf("unsupported interpolation mode: %d", mode)
	}
	inverse, err := matrix.Inverse()
	if err != nil {
		return nil, err
	}
	transformed, err := matrix.Bounds(math3d.Rect{Width: float64(source.Bounds().Dx()), Height: float64(source.Bounds().Dy())})
	if err != nil {
		return nil, err
	}
	minX, minY := transformed.X, transformed.Y
	autoViewport := viewport.Empty()
	if autoViewport {
		viewport = image.Rect(0, 0, max(1, int(math.Ceil(transformed.Width))), max(1, int(math.Ceil(transformed.Height))))
	}
	destination := image.NewRGBA(viewport)
	if background == nil {
		background = color.Transparent
	}
	draw.Draw(destination, destination.Bounds(), &image.Uniform{C: background}, image.Point{}, draw.Src)
	read := rgbaReaderFor(source)
	bounds := source.Bounds()
	for y := viewport.Min.Y; y < viewport.Max.Y; y++ {
		for x := viewport.Min.X; x < viewport.Max.X; x++ {
			destinationX, destinationY := float64(x)+0.5, float64(y)+0.5
			if autoViewport {
				destinationX = minX + float64(x-viewport.Min.X) + 0.5
				destinationY = minY + float64(y-viewport.Min.Y) + 0.5
			}
			point := inverse.Apply(math3d.Point2{X: destinationX, Y: destinationY})
			if sampled, ok := sample(read, bounds, point.X-0.5, point.Y-0.5, mode); ok {
				// Write the premultiplied bytes directly: destination.Set boxes
				// the value in a color.Color interface, which allocates per pixel.
				offset := destination.PixOffset(x, y)
				destination.Pix[offset] = sampled.R
				destination.Pix[offset+1] = sampled.G
				destination.Pix[offset+2] = sampled.B
				destination.Pix[offset+3] = sampled.A
			}
		}
	}
	return destination, nil
}

// rgbaReader reads one source pixel as a premultiplied RGBA value.
type rgbaReader func(x, y int) color.RGBA

// rgbaReaderFor picks a direct Pix reader for the concrete image types that
// rendering produces. Indexing Pix avoids the per-pixel color.Color interface
// boxing that image.Image.At incurs; other types keep the generic path.
func rgbaReaderFor(source image.Image) rgbaReader {
	switch img := source.(type) {
	case *image.Gray:
		return func(x, y int) color.RGBA {
			value := img.Pix[img.PixOffset(x, y)]
			return color.RGBA{R: value, G: value, B: value, A: 255}
		}
	case *image.Gray16:
		return func(x, y int) color.RGBA {
			// Gray16 pixels are big-endian; RGBAModel keeps their high byte.
			value := img.Pix[img.PixOffset(x, y)]
			return color.RGBA{R: value, G: value, B: value, A: 255}
		}
	case *image.RGBA:
		return func(x, y int) color.RGBA {
			offset := img.PixOffset(x, y)
			return color.RGBA{R: img.Pix[offset], G: img.Pix[offset+1], B: img.Pix[offset+2], A: img.Pix[offset+3]}
		}
	case *image.NRGBA:
		return func(x, y int) color.RGBA {
			offset := img.PixOffset(x, y)
			return premultipliedRGBA(img.Pix[offset], img.Pix[offset+1], img.Pix[offset+2], img.Pix[offset+3])
		}
	}
	return func(x, y int) color.RGBA {
		return color.RGBAModel.Convert(source.At(x, y)).(color.RGBA)
	}
}

// premultipliedRGBA converts straight-alpha NRGBA bytes to the premultiplied
// RGBA value color.RGBAModel produces for the same pixel.
func premultipliedRGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{
		R: premultiplyChannel(r, a),
		G: premultiplyChannel(g, a),
		B: premultiplyChannel(b, a),
		A: a,
	}
}

// premultiplyChannel matches the 16-bit multiply and shift color.NRGBA.RGBA
// performs before RGBAModel truncates the result to 8 bits.
func premultiplyChannel(value, alpha uint8) uint8 {
	if alpha == 255 {
		return value
	}
	return uint8(((uint32(value) * 0x101 * uint32(alpha)) / 0xff) >> 8)
}

func sample(read rgbaReader, bounds image.Rectangle, x, y float64, mode interpolation.Mode) (color.RGBA, bool) {
	if mode == interpolation.ModeNearestNeighbor {
		px, py := int(math.Round(x)), int(math.Round(y))
		if !image.Pt(px, py).In(bounds) {
			return color.RGBA{}, false
		}
		return read(px, py), true
	}
	if x < float64(bounds.Min.X) || y < float64(bounds.Min.Y) || x > float64(bounds.Max.X-1) || y > float64(bounds.Max.Y-1) {
		return color.RGBA{}, false
	}
	x0, y0 := int(math.Floor(x)), int(math.Floor(y))
	x1, y1 := min(x0+1, bounds.Max.X-1), min(y0+1, bounds.Max.Y-1)
	tx, ty := x-float64(x0), y-float64(y0)
	a, b := read(x0, y0), read(x1, y0)
	c, d := read(x0, y1), read(x1, y1)
	return color.RGBA{R: blend(a.R, b.R, c.R, d.R, tx, ty), G: blend(a.G, b.G, c.G, d.G, tx, ty), B: blend(a.B, b.B, c.B, d.B, tx, ty), A: blend(a.A, b.A, c.A, d.A, tx, ty)}, true
}

func blend(a, b, c, d uint8, tx, ty float64) uint8 {
	top := float64(a)*(1-tx) + float64(b)*tx
	bottom := float64(c)*(1-tx) + float64(d)*tx
	return uint8(math.Round(top*(1-ty) + bottom*ty))
}
