package minify

import (
	"github.com/h2non/bimg"
)

func ImgMinify(i []byte, q int) ([]byte, error) {
	options := bimg.Options{
		Quality: q,
		Type:    bimg.WEBP,
	}

	newImg, err := bimg.NewImage(i).Process(options)
	if err != nil {
		return nil, err
	}

	return newImg, nil
}
