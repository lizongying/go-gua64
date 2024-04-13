package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/lizongying/go-gua64/gua64"
	"os"
)

// go run cmd/gua64/main.go
func main() {
	ePtr := flag.String("e", "", "Encode String")
	dPtr := flag.String("d", "", "Decode String")
	fPtr := flag.Bool("f", false, "Input Is A File")
	oPtr := flag.String("o", "", "Output File")

	flag.Parse()
	e := *ePtr
	d := *dPtr
	f := *fPtr
	o := *oPtr

	if e != "" && d != "" {
		fmt.Println("Error: Please provide only one of -e or -d flags")
		os.Exit(1)
	}

	if e == "" && d == "" {
		fmt.Println("Error: Please provide either -e or -d flag")
		os.Exit(1)
	}

	g := gua64.NewGua64()
	if e != "" {
		in := []byte(e)
		if f {
			in, _ = os.ReadFile(e + d)
		}
		r := g.Encode(in)
		if o != "" {
			_ = os.WriteFile(o, []byte(r), 0644)
		} else {
			fmt.Println(r)
		}
	}
	if d != "" {
		if f {
			in, _ := os.ReadFile(e + d)
			d = string(bytes.TrimSpace(in))
		}
		r := g.Decode(d)
		if o != "" {
			_ = os.WriteFile(o, r, 0644)
		} else {
			fmt.Println(string(r))
		}
	}
}
