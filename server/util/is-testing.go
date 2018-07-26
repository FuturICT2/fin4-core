package util

// IsTesting returns true if app is run with env var TEST=1
func IsTesting() bool {
	return Getenv("TEST") == "1"
}
