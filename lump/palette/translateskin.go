package palette

import (
	"errors"
	"fmt"
	"image/color"
)

func TranslatePalette(skin uint, palette color.Palette) (color.Palette, error) {
  return TranslatePaletteWithStartTransColor(skin, palette, 160)
}

func TranslatePaletteWithStartTransColor(skin uint, palette color.Palette, startTransColor uint8) (color.Palette, error) {
	if skin >= MAXTRANSLATIONS {
		return palette, errors.New(fmt.Sprintf("Invalid skin number %d, see SKINCOLOR_NONE", skin))
	}

  if startTransColor - SKIN_TRANSLATION_SIZE > PALETTE_SIZE - 1 {
		return palette, errors.New(fmt.Sprintf("Invalid skin number %d, see SKINCOLOR_NONE", skin))
  }

  for i := uint8(0) ; i < SKIN_TRANSLATION_SIZE ; i++ {
    palette[startTransColor + i] = palette[skinTranslations[skin][i]]
  }
	return palette, nil
}
