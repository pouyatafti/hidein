package lib

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const pixperbyte = 3

func Encode(typ string, rd io.Reader, wt io.Writer, bytes []uint8) error {
	var in image.Image
	var err error
	switch typ {
	case "png":
		in, err = png.Decode(rd)
	case "jpeg":
		in, err = jpeg.Decode(rd)
	case "gif":
		in, err = gif.Decode(rd)
	default:
		in, _, err = image.Decode(rd)
	}
	if err != nil {
		return err
	}

	bounds := in.Bounds()

	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	unusedrows := height % pixperbyte

	maxbits := width * (height - unusedrows) * 8 / pixperbyte
	if len(bytes)*8 > maxbits {
		return errors.New("image too small")
	}

	out := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	draw.Draw(out, bounds, in, bounds.Min, draw.Src)

	i := 0
	var R, G, B, A uint32
	var r, g, b, a, u uint8
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y-unusedrows; {
			if i < len(bytes) {
				u = bytes[i]
				i++
			} else {
				u = 0
			}
			R, G, B, A = out.At(x, y).RGBA()
			y++
			r, g, b, a = uint8(R&0xFE), uint8(G&0xFE), uint8(B&0xFE), uint8(A&0xFE)
			r |= (u & (1 << 0)) >> 0
			g |= (u & (1 << 1)) >> 1
			b |= (u & (1 << 2)) >> 2
			_ = a

			out.Set(x, y-1, color.NRGBA{R: r, G: g, B: b, A: 255})

			R, G, B, A = out.At(x, y).RGBA()
			y++
			r, g, b, a = uint8(R&0xFE), uint8(G&0xFE), uint8(B&0xFE), uint8(A&0xFE)
			r |= (u & (1 << 3)) >> 3
			g |= (u & (1 << 4)) >> 4
			b |= (u & (1 << 5)) >> 5
			_ = a

			out.Set(x, y-1, color.NRGBA{R: r, G: g, B: b, A: 255})

			R, G, B, A = out.At(x, y).RGBA()
			y++
			r, g, b, a = uint8(R&0xFE), uint8(G&0xFE), uint8(B&0xFE), uint8(A&0xFE)
			r |= (u & (1 << 6)) >> 6
			g |= (u & (1 << 7)) >> 7
			_ = b
			_ = a

			out.Set(x, y-1, color.NRGBA{R: r, G: g, B: b, A: 255})
		}
	}

	return png.Encode(wt, out)
}

func Decode(rd io.Reader, bytes []uint8, l int) error {
	in, err := png.Decode(rd)
	if err != nil {
		return err
	}

	bounds := in.Bounds()

	height := bounds.Max.Y - bounds.Min.Y
	unusedrows := height % pixperbyte

	i := 0
	var R, G, B, A uint32
	var r, g, b, a, u uint8
Loop:
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y-unusedrows; {
			if i >= len(bytes) || i >= l {
				break Loop
			}
			u = 0
			R, G, B, A = in.At(x, y).RGBA()
			y++
			// r,g,b,a = uint8(R & 0xFF), uint8(G & 0xFF), uint8(B & 0xFF), uint8(A & 0xFF)
			r, g, b, a = uint8(R/257), uint8(G/257), uint8(B/257), uint8(A/257)
			u |= (r & 1) << 0
			u |= (g & 1) << 1
			u |= (b & 1) << 2
			_ = a

			R, G, B, A = in.At(x, y).RGBA()
			y++
			// r,g,b,a = uint8(R & 0xFF), uint8(G & 0xFF), uint8(B & 0xFF), uint8(A & 0xFF)
			r, g, b, a = uint8(R/257), uint8(G/257), uint8(B/257), uint8(A/257)
			u |= (r & 1) << 3
			u |= (g & 1) << 4
			u |= (b & 1) << 5
			_ = a

			R, G, B, A = in.At(x, y).RGBA()
			y++
			// r,g,b,a = uint8(R & 0xFF), uint8(G & 0xFF), uint8(B & 0xFF), uint8(A & 0xFF)
			r, g, b, a = uint8(R/257), uint8(G/257), uint8(B/257), uint8(A/257)
			u |= (r & 1) << 6
			u |= (g & 1) << 7
			_ = b
			_ = a

			bytes[i] = u
			i++
		}
	}

	return nil
}
