package img

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/lytics/logrus"
)

//FromBase64 parses image info from base64 string
func FromBase64(base64Str string) (data []byte, contentType string, ext string, err error) {
	b64data := base64Str[strings.IndexByte(base64Str, ',')+1:]
	data, err = base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("routes-assets/CreateTokens:1")
	}
	contentType = string(http.DetectContentType(data))
	parts := strings.Split(contentType, "/")
	ext = parts[len(parts)-1]
	return data, contentType, ext, nil
}
