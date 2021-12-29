package gua

import (
	"fmt"
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
	Gua       string
	gua64dict map[string]int
}

func (gua64 *Gua64) New() {
	gua64.Gua = gua
}

func (gua64 *Gua64) Encode(str string) (out string) {
	fmt.Println(str)
	return out
}

func (gua64 *Gua64) Decode(str string) (out string) {
	fmt.Println(str)
	return out
}
