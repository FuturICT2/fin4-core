package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

//HandleExifOrientation takes care of EXIF orientation
func HandleExifOrientation(img image.Image, data []byte) (image.Image, error) {
	xf, err := exif.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	orientation, err := xf.Get(exif.Orientation)
	if err != nil {
		return nil, err
	}
	switch orientation.String() {
	case "1":
		img = imaging.Clone(img)
	case "2":
		img = imaging.FlipV(img)
	case "3":
		img = imaging.Rotate180(img)
	case "4":
		img = imaging.Rotate180(imaging.FlipV(img))
	case "5":
		img = imaging.Rotate270(imaging.FlipV(img))
	case "6":
		img = imaging.Rotate270(img)
	case "7":
		img = imaging.Rotate90(imaging.FlipV(img))
	case "8":
		img = imaging.Rotate90(img)
	}

	return img, nil
}
