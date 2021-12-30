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
