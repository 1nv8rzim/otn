package config

import (
	"os"
	"time"
)

const (
	EXPIRATION    = 5 * time.Minute
	CLEANINTERVAL = 10 * time.Second
	URILENGTH     = 10
)

var (
	URL string
)

func init() {
	url := os.Getenv("URL")
	if url == "" {
		URL = "http://localhost"
	} else {
		URL = url
	}
}
