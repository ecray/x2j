package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	xj "github.com/basgys/goxml2json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convStream(d io.Reader) *bytes.Buffer {
	j, err := xj.Convert(d)
	check(err)

	return j
}

func main() {
	j := convStream(os.Stdin)
	fmt.Println(j)
}
