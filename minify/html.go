package minify

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func HtmlMinify(i []byte) ([]byte, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	o, err := m.Bytes("text/html", i)
	if err != nil {
		return nil, err
	}
	return o, nil
}
