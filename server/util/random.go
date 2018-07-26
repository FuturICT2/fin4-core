package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

// RandomString generates a random string
func RandomString() string {
	rnd := fmt.Sprintf("%d", rand.Int63())
	rbytes := md5.Sum([]byte(rnd))
	return fmt.Sprintf("%x", rbytes)
}
