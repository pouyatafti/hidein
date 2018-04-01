package main

import (
	"fmt"
	"github.com/pouyatafti/hidein/lib"
	"os"
	"strconv"
)

const defaultlength = 1024

func usage() {
	fmt.Println("usage: decode <image.png|.jpg.|.gif> [length]")
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		usage()
		os.Exit(1)
	}

	imfile := os.Args[1]
	var length int
	var err error
	if len(os.Args) > 2 {
		length, err = strconv.Atoi(os.Args[2])
		if err != nil {
			length = defaultlength
		}
	} else {
		length = defaultlength
	}

	bytes := make([]uint8, length)

	dst, err := os.Open(imfile)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	if err := lib.Decode(dst, bytes, length); err != nil {
		panic(err)
	}

	os.Stdout.Write(bytes[:])
}
