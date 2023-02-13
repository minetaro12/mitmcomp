package imgcomp

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"
)

func ImageCompress(i []byte, q int) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(i))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	err = webp.Encode(buffer, img, &webp.Options{Quality: float32(q), Lossless: false})
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
