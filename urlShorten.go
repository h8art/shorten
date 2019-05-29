package urlShorten

import (
	"encoding/hex"
	"fmt"
)

type Shortener interface {
	Shorten(string) string
	Resolve(string) string
}

var urls []url

type url struct {
	hex  string
	full string
}

func Shorten(s string) string {
	var tempUrl url
	src := []byte(s)
	tempUrl.hex = hex.EncodeToString(src)
	tempUrl.full = s
	urls = append(urls, tempUrl)
	return tempUrl.hex
}

func Resolve(s string) string {
	for _, n := range urls {
		if n.hex == s {
			return n.full
		}
	}
	return ""
}

func convert(u string) {
	uHex := Shorten(u)
	fmt.Println(uHex)
	fmt.Println(Resolve(uHex))
}
