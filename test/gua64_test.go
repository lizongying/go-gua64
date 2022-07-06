package test

import (
	"github.com/lizongying/go-gua64/gua"
	"testing"
)

// go test -v test/gua64_test.go

// TestGua64 go test -v test/gua64_test.go -run TestGua64
func TestGua64(t *testing.T) {
	g := gua64.Gua64{}
	g.New()
	t.Log(g.Encode("hello，世界"))
	t.Log(g.Decode("䷯䷬䷿䷶䷸䷬䷀䷌䷌䷎䷼䷲䷰䷳䷸䷘䷔䷭䷒☯"))
}
