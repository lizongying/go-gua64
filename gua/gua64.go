package gua

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
	"䷊䷙䷄䷈䷡䷍䷪䷀"

type Gua64 struct {
	gua64list []string
	gua64dict map[string]int
}

func (gua64 *Gua64) New() {
	gua64.gua64list = strings.Split(gua, "")
	gua64.gua64dict = make(map[string]int)
	for i := 0; i < 64; i++ {
		gua64.gua64dict[gua64.gua64list[i]] = i
	}
}

func (gua64 *Gua64) Encode(str string) (out string) {
	var encoded []string
	strLen := len(str)
	var begin uint8
	for i := 0; i < strLen; i += 3 {
		begin = str[i] >> 2
		encoded = append(encoded, gua64.gua64list[begin])
		if i == strLen-1 {
			begin = (str[i] & 0x3) << 4
			encoded = append(encoded, gua64.gua64list[begin])
			encoded = append(encoded, "☯")
			encoded = append(encoded, "☯")
			continue
		}
		begin = (str[i]&0x3)<<4 | str[i+1]>>4
		encoded = append(encoded, gua64.gua64list[begin])
		if i == strLen-2 {
			begin = (str[i+1] & 0xf) << 2
			encoded = append(encoded, gua64.gua64list[begin])
			encoded = append(encoded, "☯")
			continue
		}
		begin = (str[i+1]&0xf)<<2 | str[i+2]>>6
		encoded = append(encoded, gua64.gua64list[begin])
		begin = str[i+2] & 0x3f
		encoded = append(encoded, gua64.gua64list[begin])
	}
	return strings.Join(encoded, "")
}

func (gua64 *Gua64) Decode(str string) (out string) {
	var b []int
	strLen := len(str)
	for i := 0; i < strLen; i += 3 {
		v, ok := gua64.gua64dict[str[i:i+3]]
		if ok {
			b = append(b, v)
			continue
		}
		b = append(b, 0)
	}
	var encoded []byte
	outLen := len(b)
	for i := 0; i < outLen; i += 4 {
		encoded = append(encoded, byte((b[i]&0x3f)<<2|(b[i+1]>>4&0x3)))
		two := (b[i+1]&0xf)<<4 | (b[i+2] >> 2 & 0xf)
		if two != 0 {
			encoded = append(encoded, byte(two))
		}
		three := (b[i+2]&0x3)<<6 | (b[i+3] & 0x3f)
		if three != 0 {
			encoded = append(encoded, byte(three))
		}
	}
	return string(encoded)
}
