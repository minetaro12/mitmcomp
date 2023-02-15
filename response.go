package main

import (
	"mitmcomp/transcoder"
	"strconv"
	"strings"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

type MinifyAddon struct {
	proxy.BaseAddon
}

func (a *MinifyAddon) Response(f *proxy.Flow) {
	originalSize := len(f.Response.Body)
	if strings.Contains(f.Response.Header.Get("Content-Type"), "image/") {
		// 画像の圧縮
		transcoder.ImgRecompress(f, *imgQuality)
		if *debugMode {
			debug(originalSize, len(f.Response.Body), "img", f)
		}
		minifyLog(originalSize, len(f.Response.Body), f.Request.URL)
	} else {
		if *brotliComp {
			// brotliで再圧縮
			if err := transcoder.Recompress(f); err != nil {
				return
			}
			if *debugMode {
				debug(originalSize, len(f.Response.Body), "brotli", f)
			}
			minifyLog(originalSize, len(f.Response.Body), f.Request.URL)
		} else {
			return
		}
	}

	// ヘッダーの書き換え
	f.Response.Header.Del("Content-Length")
	f.Response.Header.Add("Content-Length", strconv.Itoa(len(f.Response.Body)))
}
