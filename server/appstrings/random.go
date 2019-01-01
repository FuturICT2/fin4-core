package appstrings

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomString generates a random string
func GetRandomString(length int) string {
	rnd := rand.Int63()
	rbytes := md5.Sum([]byte(fmt.Sprintf("%d", rnd)))
	s := fmt.Sprintf("%x", rbytes)
	if length != 0 {
		return s[0:length]
	}
	return s
}
