package main

import (
	"io"
	"os"
)

func main() {
	file := string(os.Args[1])
	data, _ := os.Open(file)
	io.Copy(os.Stdout, data)
}
