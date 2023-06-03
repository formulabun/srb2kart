package graphic

import (
	"image"
	"image/color"
)

func graphicToImage(h *header, p color.Palette) (image.Image, error) {
	rect := image.Rectangle{
		image.Point{int(h.leftOffset), int(h.topOffset)},
		image.Point{int(h.leftOffset) + int(h.width), int(h.topOffset) + int(h.height)},
	}
	result := image.NewPaletted(rect, p)

	for c, post := range h.posts {
		for o, pixel := range post.data {
			r := o + int(post.topDelta)
			result.SetColorIndex(c, r, pixel)
		}
	}

	return result, nil
}
