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

		// gzipを展開
		raw, err := gzipDecompress(f.Response.Body)
		if err != nil {
			return err
		}

		// brotliで圧縮
		br, err := brotliCompress(raw)
		if err != nil {
			return err
		}

		f.Response.Body = br

		// ヘッダーの書き換え
		f.Response.Header.Del("Content-Encoding")
		f.Response.Header.Add("Content-Encoding", "br")
		return nil

	} else if f.Response.Header.Get("Content-Encoding") == "" && strings.Contains(f.Request.Header.Get("Accept-Encoding"), "br") {

		// brotliで圧縮
		br, err := brotliCompress(f.Response.Body)
		if err != nil {
			return err
		}

		f.Response.Body = br

		// ヘッダーの書き換え
		f.Response.Header.Del("Content-Encoding")
		f.Response.Header.Add("Content-Encoding", "br")
		return nil

	} else {
		return errors.New("can't compress")
	}
}

func gzipDecompress(i []byte) ([]byte, error) {
	gzReader, err := gzip.NewReader(bytes.NewReader(i))
	if err != nil {
		return nil, err
	}
	var rawBuffer = new(bytes.Buffer)
	io.Copy(rawBuffer, gzReader)

	return rawBuffer.Bytes(), nil
}

func brotliCompress(i []byte) ([]byte, error) {
	var brBuffer = new(bytes.Buffer)
	brWriter := brotli.NewWriter(brBuffer)
	brWriter.Write(i)
	if err := brWriter.Close(); err != nil {
		return nil, err
	}

	return brBuffer.Bytes(), nil
}
