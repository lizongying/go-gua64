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
	ePtr := flag.String("e", "", "Input the string to be encoded")
	dPtr := flag.String("d", "", "Input the string to be decoded")
	vPtr := flag.String("v", "", "Input the string and judge whether it is gua64")
	fPtr := flag.Bool("f", false, "Indicate whether the input is a file")
	oPtr := flag.String("o", "", "Specify the path of the output file")
	helpPtr := flag.Bool("h", false, "Show help information")

	flag.Parse()
	e := *ePtr
	d := *dPtr
	v := *vPtr
	f := *fPtr
	o := *oPtr

	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	if e == "" && d == "" && v == "" {
		fmt.Println("Error: Please provide either -e or -d or -v flag")
		os.Exit(1)
	}

	if e == "" && d == "" && v == "" {
		fmt.Println("Error: Please provide either -e or -d or -v flag. Use -h for more information.")
		os.Exit(1)
	}

	count := 0
	if e != "" {
		count++
	}
	if d != "" {
		count++
	}
	if v != "" {
		count++
	}
	if count > 1 {
		fmt.Println("Error: Please provide only one of -e or -d or -v flag. Use -h for more information.")
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
			err := os.WriteFile(o, []byte(r), 0644)
			if err != nil {
				fmt.Printf("Error: Failed to write to file %s. %v\n", o, err)
				os.Exit(1)
			}
		} else {
			fmt.Println(r)
		}
	}
	if d != "" {
		if f {
			in, err := os.ReadFile(e + d)
			if err != nil {
				fmt.Printf("Error: Failed to read file %s. %v\n", o, err)
				os.Exit(1)
			}
			d = string(bytes.TrimSpace(in))
		}
		r := g.Decode(d)
		if o != "" {
			err := os.WriteFile(o, r, 0644)
			if err != nil {
				fmt.Printf("Error: Failed to write to file %s. %v\n", o, err)
				os.Exit(1)
			}
		} else {
			fmt.Println(string(r))
		}
	}
	if v != "" {
		if f {
			in, err := os.ReadFile(e + d)
			if err != nil {
				fmt.Printf("Error: Failed to read file %s. %v\n", o, err)
				os.Exit(1)
			}
			d = string(bytes.TrimSpace(in))
		}
		r := g.Verify(d)
		fmt.Println(r)
	}
}
