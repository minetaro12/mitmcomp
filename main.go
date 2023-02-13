package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"mitmcomp/imgcomp"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

var (
	bindAddr   = flag.String("b", "0.0.0.0", "bind address")
	bindPort   = flag.Int("p", 8080, "bind port")
	caPath     = flag.String("ca", "./ca", "ca certificates")
	imgQuality = flag.Int("q", 10, "webp quality")
)

type ImgCompress struct {
	proxy.BaseAddon
}

func (a *ImgCompress) Response(f *proxy.Flow) {
	// 画像でない場合はそのまま返す
	if !strings.Contains(f.Response.Header.Get("Content-Type"), "image/") {
		return
	}

	// 画像を圧縮
	img, err := imgcomp.ImageCompress(f.Response.Body, *imgQuality)
	if err != nil {
		return
	}
	f.Response.Body = img

	// ログ
	bs, _ := strconv.Atoi(f.Response.Header.Get("Content-Length"))
	as := len(f.Response.Body)
	p := (float32(as) / float32(bs)) * 100
	log.Printf("%v → %v (%v%%) %v\n", bs, as, p, f.Request.URL)

	// ヘッダーの書き換え
	f.Response.Header.Del("Content-Length")
	f.Response.Header.Del("Content-Type")
	f.Response.Header.Add("Content-Length", strconv.Itoa(as))
	f.Response.Header.Add("Content-Type", "image/webp")
}

func main() {
	flag.Parse()
	opts := &proxy.Options{
		Addr:              fmt.Sprintf("%v:%v", *bindAddr, *bindPort),
		CaRootPath:        *caPath,
		StreamLargeBodies: 1024 * 1024 * 5,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		log.Fatal(err)
	}

	p.AddAddon(&ImgCompress{})

	log.Fatal(p.Start())
}
