package xtract

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

var unlim = -1
var trimFunc = unicode.IsSpace

//TrimFunc trim function applied in resule
type TrimFunc func(r rune) bool

//SetTrimFunc set trim function, default is trim left/right spaces with unicode.IsSpace
func SetTrimFunc(f TrimFunc) {
	trimFunc = f
}

// Page extract text from the given url
func Page(url string) (string, error) {
	return PageLim(url, unlim)
}

// PageLim extract text from the page and return upto lim number of words
func PageLim(url string, lim int) (string, error) {
	rs, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer rs.Body.Close()
	return xtract(rs.Body, lim), err
}

//Value extract text from the given html value
func Value(htmlVal string) string {
	return ValueLim(htmlVal, unlim)
}

//ValueLim extract text from the given html value and return upto lim number of words
func ValueLim(htmlVal string, lim int) string {
	b := bytes.NewReader([]byte(htmlVal))
	return xtract(b, lim)
}

func xtract(r io.Reader, lim int) string {
	z := html.NewTokenizer(r)
	rs := bytes.NewBufferString("")

	for {
		t := z.Next()
		if t == html.ErrorToken {
			return rs.String()
		}
		if t == html.TextToken {
			if trimFunc != nil {
				if rs.Len() > 0 {
					rs.Write([]byte(" "))
				}
				rs.Write(bytes.TrimFunc(z.Text(), trimFunc))
			} else {
				rs.Write(z.Text())
			}
			if lim != unlim {
				v := strings.Fields(rs.String())
				wc := len(v)
				if wc >= lim {
					return strings.Join(v[:min(wc, lim)], " ")
				}
			}
		}
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
