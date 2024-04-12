package gua64

import (
	"testing"
)

// go test -v gua64/*.go

// TestGua64 go test -v gua68/*.go -run TestGua64
func TestGua64(t *testing.T) {
	g := NewGua64()
	t.Log(g.Encode(nil))
	t.Log(g.Encode([]byte("")))
	t.Log(g.Encode([]byte("l")))
	t.Log(string(g.Decode("䷸⚏")))
	t.Log(g.Encode([]byte("l0")))
	t.Log(string(g.Decode("䷸䷓⚏⚏")))
}
