package img

import (
	"bytes"
	"encoding/base64"
	"image"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/lytics/logrus"
)

var imgMaxWidth = 900

//FromBase64 parses image info from base64 string
func FromBase64(base64Str string) (data []byte, contentType string, ext string, err error) {
	b64data := base64Str[strings.IndexByte(base64Str, ',')+1:]
	data, err = base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("img/img:1")
		return nil, "", "", err
	}
	contentType = string(http.DetectContentType(data))
	parts := strings.Split(contentType, "/")
	ext = strings.ToLower(parts[len(parts)-1])

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, "", "", err
	}

	if imgxif, err1 := HandleExifOrientation(img, data); err1 == nil {
		img = imgxif
	}

	b := img.Bounds()
	if b.Max.X > imgMaxWidth {
		img = imaging.Resize(img, imgMaxWidth, 0, imaging.Lanczos)
	}

	var buf bytes.Buffer
	if ext == "jpg" || ext == "jpeg" {
		err = imaging.Encode(&buf, img, imaging.JPEG)
	} else if ext == "png" {
		err = imaging.Encode(&buf, img, imaging.PNG)
	} else if ext == "gif" {
		err = imaging.Encode(&buf, img, imaging.GIF)
	}
	if err != nil {
		return nil, "", "", err
	}
	return buf.Bytes(), contentType, ext, nil
}
