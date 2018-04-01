package main

import (
	"fmt"
	"github.com/pouyatafti/hidein/lib"
	"io/ioutil"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Println("usage: encode <image.png|.jpg.|.gif> <data|->")
}

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	imfile := os.Args[1]
	dtafile := os.Args[2]

	typ := filepath.Ext(imfile)[1:]

	src, err := os.Open(imfile)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	var dta []uint8
	if dtafile == "-" {
		dta, err = ioutil.ReadAll(os.Stdin)
	} else {
		dta, err = ioutil.ReadFile(dtafile)
	}
	if err != nil {
		panic(err)
	}

	dst := os.Stdout

	if err := lib.Encode(typ, src, dst, dta); err != nil {
		panic(err)
	}
}
