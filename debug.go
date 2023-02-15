package main

import (
	"fmt"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

func debug(bs int, as int, m string, f *proxy.Flow) {
	p := (float32(as) / float32(bs)) * 100
	str := fmt.Sprintf("%v%%", p)
	f.Response.Header.Add("Mitmcomp", m)
	f.Response.Header.Add("Mitmcomp-Debug", str)
}
