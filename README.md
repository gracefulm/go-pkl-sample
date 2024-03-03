# GoでPklを使うサンプル

[この記事](https://qiita.com/gracefulm/items/ed5b7f0054b4cd9f63d4)でPklを紹介した時のサンプルプログラムです。

## How to use

※Pklがインストールされている必要があります。

```sh
git clone https://github.com/gracefulm/go-pkl-sample.git
cd go-pkl-sample
go mod tidy
go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest
go run main.go
```
