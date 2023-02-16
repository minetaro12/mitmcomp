# mitmcomp
HTTP,HTTPSのコンテンツを圧縮するフォワードプロキシです。  
デフォルトですべての画像をwebp クオリティ10で圧縮します。  
端末に`mitmproxy-ca-cert.pem`をインストールする必要があります。(デフォルト設定では`./ca`内にあります。)

```
#ビルド
$ go build

#実行
$ ./mitmcomp

#ポートを指定して実行
$ ./mitmcomp -p=8000

#brotli圧縮を有効にする
$ ./mitmcomp -br=true
```

## Docker
```
$ docker run -it --rm -v ./ca:/app/ca -p 8080:8080 ghcr.io/minetaro12/mitmcomp

#docker compose
$ docker compose up -d
```

## Thanks
- https://github.com/lqqyt2423/go-mitmproxy
- https://github.com/andybalholm/brotli
- https://github.com/h2non/bimg