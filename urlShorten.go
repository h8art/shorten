package urlShorten

import (
	"encoding/hex"
	"fmt"
)

type MyShortener struct {
	storage map[string]string
}

type Shortener interface {
	Shorten(string) string
	Resolve(string) string
}

func (s MyShortener) Shorten(url string) string {
	src := []byte(url)
	s.storage[hex.EncodeToString(src)] = url
	return hex.EncodeToString(src)
}

func (s MyShortener) Resolve(url string) string {
	return s.storage[url]
}

func convert(u string) {
	tempUrls := new(MyShortener)
	tempUrls.storage = make(map[string]string)

	urlShorten := Shortener.Shorten(tempUrls, u)
	fmt.Println(urlShorten)

	urlResoled := Shortener.Resolve(tempUrls, urlShorten)
	fmt.Println(urlResoled)
}
