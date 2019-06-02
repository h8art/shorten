package urlShorten

import (
	"encoding/hex"
	"fmt"
)

type MyShortened struct {
	storage map[string]string
}

type Shortened interface {
	Shorten(string) string
	Resolve(string) string
}

func (s MyShortened) Shorten(url string) string {
	src := []byte(url)
	s.storage[hex.EncodeToString(src)] = url
	return hex.EncodeToString(src)
}

func (s MyShortened) Resolve(url string) string {
	return s.storage[url]
}
func NewMyShortener() *MyShortened {
	shortener := new(MyShortened)
	shortener.storage = make(map[string]string)
	return shortener
}
func Convert(u string) {
	tempUrls := NewMyShortener()
	urlShorten := tempUrls.Shorten(u)

	fmt.Println(urlShorten)

	urlResoled := tempUrls.Resolve(urlShorten)
	fmt.Println(urlResoled)
}
