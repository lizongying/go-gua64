# 六十四卦编码

六十四卦编码，golang实现。
如：“hello，世界”会编码为“䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒〇”。

[go-gua64](https://github.com/lizongying/go-gua64)

## All language

* [golang](https://github.com/lizongying/go-gua64)
* [js](https://github.com/lizongying/js-gua64)
* [java](https://github.com/lizongying/java-gua64)
* [php](https://github.com/lizongying/php-gua64)
* [python](https://github.com/lizongying/pygua64)
* [c#](https://github.com/lizongying/dotnet-gua64)

## Download

[releases](https://github.com/lizongying/go-gua64/releases/latest)

## Usage

* -e Input the string to be encoded
    ```shell
    ./gua64_darwin_arm64 -e 123   
    # output: ䷽䷺䷎䷼
    ```
* -d Input the string to be decoded
    ```shell
    ./gua64_darwin_arm64 -d ䷽䷺䷎䷼   
    # output: 123
    ```
* -v Input the string and judge whether it is gua64
    ```shell
    ./gua64_darwin_arm64 -v ䷽䷺䷎䷼   
    # output: true
    ```
* -o Specify the path of the output file
    ```shell
    ./gua64_darwin_arm64 -e 123 -o encode.txt   
    ```

* -f Indicate whether the input is a file
    ```shell
    ./gua64_darwin_arm64 -f -d encode.txt -o decode.txt  
  
    ./gua64_darwin_arm64 -f -d decode.txt -o encode.txt    
    ```

## Build

```shell
make
```

## Docker image

[ghcr.io](https://github.com/lizongying/go-gua64/pkgs/container/go-gua64)

[hub.docker.com](https://hub.docker.com/r/lizongying/go-gua64)

## Some sample in Golang

[sample](./sample)

### Document

[document](https://pkg.go.dev/github.com/lizongying/go-gua64)

### Install

```
go get github.com/lizongying/go-gua64
```

```go
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
	fmt.Println(g.Verify(`䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒〇`))
}
```