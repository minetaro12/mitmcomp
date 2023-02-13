package transcoder

import (
	"mitmcomp/minify"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

func ImgRecompress(f *proxy.Flow, q int) {
	// 画像の圧縮
	out, err := minify.ImgMinify(f.Response.Body, q)
	if err != nil {
		return
	}

	// ヘッダーの書き換え
	f.Response.Header.Del("Content-Type")
	f.Response.Header.Add("Content-Type", "image/webp")
	f.Response.Body = out
}
