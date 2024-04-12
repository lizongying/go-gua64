# 六十四卦

六十四卦编码，golang实现。
如：“hello，世界”会编码为“䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒〇”。

[go-gua64](https://github.com/lizongying/go-gua64)

## others

* [js-gua64](https://github.com/lizongying/js-gua64)
* [php-gua64](https://github.com/lizongying/php-gua64)
* [pygua64](https://github.com/lizongying/pygua64)

## usage

* -e encode
    ```shell
    ./gua64_darwin_arm64 -e 123   
    # output: ䷽䷺䷎䷼
    ```
* -d decode
    ```shell
    ./gua64_darwin_arm64 -d ䷽䷺䷎䷼   
    # output: 123
    ```

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
	fmt.Println(string(g.Decode(`䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒〇`)))
}
```