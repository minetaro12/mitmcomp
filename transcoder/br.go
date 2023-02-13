package transcoder

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

func Recompress(f *proxy.Flow) error {
	if f.Response.Header.Get("Content-Encoding") == "gzip" && strings.Contains(f.Request.Header.Get("Accept-Encoding"), "br") {
		var rawBuffer = new(bytes.Buffer)

		// gzipを展開
		gzReader, err := gzip.NewReader(bytes.NewReader(f.Response.Body))
		if err != nil {
			return err
		}

		io.Copy(rawBuffer, gzReader)

		// brotliで圧縮
		var brBuffer bytes.Buffer
		brWriter := brotli.NewWriter(&brBuffer)
		brWriter.Write(rawBuffer.Bytes())
		if err = brWriter.Close(); err != nil {
			return err
		}

		f.Response.Body = brBuffer.Bytes()

		// ヘッダーの書き換え
		f.Response.Header.Del("Content-Encoding")
		f.Response.Header.Add("Content-Encoding", "br")
		return nil

	} else if f.Response.Header.Get("Content-Encoding") == "" && strings.Contains(f.Request.Header.Get("Accept-Encoding"), "br") {

		// brotliで圧縮
		var brBuffer bytes.Buffer
		brWriter := brotli.NewWriter(&brBuffer)
		brWriter.Write(f.Response.Body)
		if err := brWriter.Close(); err != nil {
			return err
		}

		f.Response.Body = brBuffer.Bytes()

		// ヘッダーの書き換え
		f.Response.Header.Del("Content-Encoding")
		f.Response.Header.Add("Content-Encoding", "br")
		return nil

	} else {
		return errors.New("can't compress")
	}
}
