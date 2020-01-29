package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	xj "github.com/basgys/goxml2json"
)

func convStream(d io.Reader) (*bytes.Buffer, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error converting document.")
		}
	}()
	j, err := xj.Convert(d, xj.WithAttrPrefix(""))

	return j, err

}

func main() {
	j, err := convStream(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(j)
}
