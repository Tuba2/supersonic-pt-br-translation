package theme

import (
	"errors"
	"image/color"
	"io/ioutil"
	"log"
	"strings"
	"supersonic/res"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const ColorNamePageBackground fyne.ThemeColorName = "PageBackground"

var (
	normalFont fyne.Resource
	boldFont   fyne.Resource
)

type MyTheme struct {
	NormalFont string
	BoldFont   string
}

const (
	IconNameNowPlaying  fyne.ThemeIconName = "NowPlaying"
	IconNameFavorite    fyne.ThemeIconName = "Favorite"
	IconNameNotFavorite fyne.ThemeIconName = "NotFavorite"
	IconNameAlbum       fyne.ThemeIconName = "Album"
	IconNameArtist      fyne.ThemeIconName = "Artist"
	IconNameGenre       fyne.ThemeIconName = "Genre"
	IconNamePlaylist    fyne.ThemeIconName = "Playlist"
	IconNameShuffle     fyne.ThemeIconName = "Shuffle"
)

var _ fyne.Theme = (*MyTheme)(nil)

func (m *MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case ColorNamePageBackground:
		if variant == theme.VariantDark {
			return color.RGBA{R: 15, G: 15, B: 15, A: 255}
		}
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case theme.ColorNameBackground:
		if variant == theme.VariantDark {
			return color.RGBA{R: 35, G: 35, B: 35, A: 255}
		}
		return color.RGBA{R: 240, G: 240, B: 240, A: 255}
	case theme.ColorNameScrollBar:
		if variant == theme.VariantDark {
			return theme.DarkTheme().Color(theme.ColorNameForeground, variant)
		}
		return theme.LightTheme().Color(theme.ColorNameForeground, variant)
	case theme.ColorNameButton:
		if variant == theme.VariantDark {
			return color.RGBA{R: 20, G: 20, B: 20, A: 50}
		}
		return color.RGBA{R: 200, G: 200, B: 200, A: 240}
	case theme.ColorNameInputBackground:
		if variant == theme.VariantDark {
			return color.RGBA{R: 20, G: 20, B: 20, A: 50}
		}
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return color.RGBA{R: 10, G: 10, B: 10, A: 255}
		}
	case theme.ColorNamePrimary:
		if variant == theme.VariantLight {
			return color.RGBA{R: 25, G: 25, B: 250, A: 255}
		}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m *MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	switch name {
	case IconNameAlbum:
		return res.ResDiscInvertPng
	case IconNameArtist:
		return res.ResPeopleInvertPng
	case IconNameFavorite:
		return res.ResHeartFilledInvertPng
	case IconNameNotFavorite:
		return res.ResHeartOutlineInvertPng
	case IconNameGenre:
		return res.ResTheatermasksInvertPng
	case IconNameNowPlaying:
		return res.ResHeadphonesInvertPng
	case IconNamePlaylist:
		return res.ResPlaylistInvertPng
	case IconNameShuffle:
		return res.ResShuffleInvertSvg
	default:
		return theme.DefaultTheme().Icon(name)
	}
}

func (m *MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	switch style {
	case fyne.TextStyle{}:
		if m.NormalFont != "" && normalFont == nil {
			if content, err := readTTFFile(m.NormalFont); err != nil {
				m.NormalFont = ""
				m.BoldFont = ""
			} else {
				normalFont = fyne.NewStaticResource("normalFont", content)
			}
		}
		if normalFont != nil {
			return normalFont
		}
	case fyne.TextStyle{Bold: true}:
		if m.BoldFont != "" && boldFont == nil {
			if content, err := ioutil.ReadFile(m.BoldFont); err != nil {
				m.BoldFont = ""
			} else {
				normalFont = fyne.NewStaticResource("boldFont", content)
			}
		}
		if boldFont != nil {
			return boldFont
		}
		if normalFont != nil {
			return normalFont
		}
	}

	return theme.DefaultTheme().Font(style)
}

func (m *MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func readTTFFile(filepath string) ([]byte, error) {
	if !strings.HasSuffix(filepath, ".ttf") {
		err := errors.New("only .ttf fonts are supported")
		log.Printf("error loading custom font %q: %s", filepath, err.Error())
		return nil, err
	}
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("error loading custom font %q: %s", filepath, err.Error())
	}
	return content, err
}
