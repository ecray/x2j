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

/*
func fileRW(inFile string, outFile string) {

	if _, err := os.Stat(inFile); os.IsNotExist(err) {
		fmt.Println("Input file not found!")
		panic(err)
	}

	d, err := os.OpenFile(inFile, 0, 0644)
	check(err)

	defer d.Close()

	x, err := d.Bytes()
	j := convStream(x)

	err = ioutil.WriteFile(outFile, j.Bytes(), 0644)
	check(err)
}
*/

func main() {
	j := convStream(os.Stdin)
	fmt.Println(j)
}
