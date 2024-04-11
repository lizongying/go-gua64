package main

import (
	"flag"
	"fmt"
	"github.com/lizongying/go-gua64/gua64"
	"os"
)

// go run cmd/gua64/main.go
func main() {
	ePtr := flag.String("e", "", "Encode String")
	dPtr := flag.String("d", "", "Decode String")
	flag.Parse()
	e := *ePtr
	d := *dPtr

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
		fmt.Println(g.Encode([]byte(e)))
	}
	if d != "" {
		fmt.Println(string(g.Decode(d)))
	}
}
