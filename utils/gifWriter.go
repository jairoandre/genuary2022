package utils

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

type GifWriter struct {
	File       *os.File
	GIF        *gif.GIF
	FileName   string
	Recording  bool
	FrameCount int
	StopCount  int
}

func RecordGif(img image.Image, writer *GifWriter) error {
	if !writer.Recording {
		return nil
	}
	if writer.FrameCount > writer.StopCount {
		err := gif.EncodeAll(writer.File, writer.GIF)
		defer writer.File.Close()
		if err != nil {
			return err
		}
		writer.Recording = false
		return nil
	}
	if writer.File == nil {
		writer.GIF = &gif.GIF{}
		file, err := os.Create(writer.FileName)
		if err != nil {
			writer.Recording = false
			return err
		}
		writer.File = file
	}
	pImage := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(pImage, pImage.Rect, img, img.Bounds().Min, draw.Over)
	writer.GIF.Image = append(writer.GIF.Image, pImage)
	writer.GIF.Delay = append(writer.GIF.Delay, 0)
	writer.FrameCount += 1
	return nil
}
