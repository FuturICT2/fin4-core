package util

import "strings"

var baseURL string

func init() {
	baseURL = MustGetenv("BASE_URL")
	baseURL = strings.TrimSpace(baseURL)
	baseURL = strings.TrimRight(baseURL, "/")
}

// URL returns a full URL
func URL(fragment string) string {
	if !strings.HasPrefix(fragment, "/") {
		fragment = "/" + fragment
	}
	return baseURL + fragment
}
