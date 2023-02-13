package main

import (
	"log"
	"net/url"
)

func minifyLog(bs int, as int, reqUrl *url.URL) {
	p := (float32(as) / float32(bs)) * 100
	log.Printf("%v → %v (%v%%) %v\n", bs, as, p, reqUrl)
}
