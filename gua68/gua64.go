package gua64

import (
	"strings"
)

const gua = "䷁䷖䷇䷓䷏䷢䷬䷋" +
	"䷎䷳䷦䷴䷽䷷䷞䷠" +
	"䷆䷃䷜䷺䷧䷿䷮䷅" +
	"䷭䷑䷯䷸䷟䷱䷛䷫" +
	"䷗䷚䷂䷩䷲䷔䷐䷘" +
	"䷣䷕䷾䷤䷶䷝䷰䷌" +
	"䷒䷨䷻䷼䷵䷥䷹䷉" +
	"䷊䷙䷄䷈䷡䷍䷪䷀" +
	"⚏⚎⚍⚌"

type Gua64 interface {
	Encode(in []byte) (out string)
	Decode(in string) (out []byte)
}
type gua64 struct {
	list []string
	dict map[string]uint8
}

func NewGua64() Gua64 {
	g := new(gua64)
	g.list = strings.Split(gua, "")
	g.dict = make(map[string]uint8)
	for i := 0; i < 68; i++ {
		g.dict[g.list[i]] = uint8(i)
	}
	return g
}

func (g *gua64) Encode(in []byte) (out string) {
	var encoded []string
	strLen := len(in)
	var begin uint8
	for i := 0; i < strLen; i += 3 {
		begin = in[i] >> 2
		encoded = append(encoded, g.list[begin])
		if i == strLen-1 {
			begin = in[i] & 0x3
			encoded = append(encoded, g.list[begin+64])
			continue
		}
		begin = (in[i]&0x3)<<4 | in[i+1]>>4
		encoded = append(encoded, g.list[begin])
		if i == strLen-2 {
			begin = (in[i+1] & 0xf) >> 2
			encoded = append(encoded, g.list[begin+64])
			begin = in[i+1] & 0x3
			encoded = append(encoded, g.list[begin+64])
			continue
		}
		begin = (in[i+1]&0xf)<<2 | in[i+2]>>6
		encoded = append(encoded, g.list[begin])
		begin = in[i+2] & 0x3f
		encoded = append(encoded, g.list[begin])
	}
	return strings.Join(encoded, "")
}

func (g *gua64) Decode(in string) (out []byte) {
	var b []uint8
	strLen := len(in)
	for i := 0; i < strLen; i += 3 {
		v, ok := g.dict[in[i:i+3]]
		if ok {
			b = append(b, v)
			continue
		}
	}
	var encoded []byte
	outLen := len(b)
	for i := 0; i < outLen; i += 4 {
		if b[i+1] >= 64 {
			encoded = append(encoded, (b[i]&0x3f)<<2|b[i+1])
			break
		}
		encoded = append(encoded, (b[i]&0x3f)<<2|(b[i+1]>>4&0x3))
		if i+2 >= outLen {
			break
		}
		if b[i+3] < 64 {
			panic("Decoding Error")
		}
		if b[i+2] >= 64 {
			two := (b[i+1]&0xf)<<4 | b[i+2]<<2 | b[i+3]
			if two != 0 {
				encoded = append(encoded, two)
			}
			break
		}
		two := (b[i+1]&0xf)<<4 | (b[i+2] >> 2 & 0xf)
		if two != 0 {
			encoded = append(encoded, two)
		}
		if b[i+3] >= 64 {
			panic("Decoding Error")
		}
		three := (b[i+2]&0x3)<<6 | (b[i+3] & 0x3f)
		if three != 0 {
			encoded = append(encoded, three)
		}
	}
	return encoded
}
