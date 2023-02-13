package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

var (
	bindAddr   = flag.String("b", "0.0.0.0", "bind address")
	bindPort   = flag.Int("p", 8080, "bind port")
	caPath     = flag.String("ca", "./ca", "ca certificates")
	imgQuality = flag.Int("q", 10, "webp quality")
	brotliComp = flag.Bool("br", false, "brotli recompression")
)

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

	p.AddAddon(&MinifyAddon{})

	log.Fatal(p.Start())
}
