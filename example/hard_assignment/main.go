package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type File interface {}

func main() {

	fileName := strings.Join((os.Args[1:]), " ")
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}