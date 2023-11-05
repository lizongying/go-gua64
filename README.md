# 六十四卦

六十四卦替代base64

## document

[document](https://pkg.go.dev/github.com/lizongying/go-gua64)

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
	fmt.Println(g.Encode([]byte("hello，世界")))
	fmt.Println(string(g.Decode("䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒☯")))
}
```