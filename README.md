# 六十四卦

六十四卦代替base64

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
	gua64 := gua.Gua64{}
	gua64.New()
	fmt.Println(gua64.Encode("hello，世界"))
	fmt.Println(gua64.Decode("䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒☯"))
}
```