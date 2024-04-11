# 六十四卦

六十四卦编码。
如：“hello，世界”会编码为“䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒☯”。

## document

[document](https://pkg.go.dev/github.com/lizongying/go-gua64)

## download

[releases](https://github.com/lizongying/go-gua64/releases/latest)

## image

[ghcr.io](https://github.com/lizongying/go-gua64/pkgs/container/go-gua64)

[hub.docker.com](https://hub.docker.com/r/lizongying/go-gua64)

## build

```shell
make
```

## install

```
go get github.com/lizongying/go-gua64
```

## sample

[sample](./sample)

```
package main

import (
	"fmt"
	"github.com/lizongying/go-gua64/gua64"
)

// go run sample/main.go
func main() {
	g := gua64.NewGua64()
	fmt.Println(g.Encode([]byte(`hello，世界`)))
	fmt.Println(string(g.Decode(`䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒☯`)))
}
```