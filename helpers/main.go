package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/1nv8rzim/otn/config"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateURI() string {
	b := make([]rune, config.URILENGTH)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateURL(uri string) string {
	return fmt.Sprintf("%s/%s\n", config.URL, uri)
}
