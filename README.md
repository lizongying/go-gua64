# 六十四卦

六十四卦代替base64

## document

[document](https://pkg.go.dev/github.com/lizongying/go-gua64)

## install

```
go get github.com/lizongying/go-gua64
```

## example

```
package main

import (
	"fmt"
	"github.com/lizongying/go-gua64/gua"
)

func main() {
	g := gua64.Gua64{}
	g.New()
	fmt.Println(g.Encode("hello，世界"))
	fmt.Println(g.Decode("䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒☯"))
}
```