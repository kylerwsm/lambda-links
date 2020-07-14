package util

import (
	"io"
	"os"
	"strings"
)

// GetNotFoundHTML fetches the not-found-page html string.
func GetNotFoundHTML() string {
	f, err := os.Open("./public/not-found.html")
	LogIf(err)
	buf := new(strings.Builder)
	io.Copy(buf, f)
	return buf.String()
}
