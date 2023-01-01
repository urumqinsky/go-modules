package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/encoder"
	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/reader"
)

func main() {
	var original reader.DBReader = encoder.NewOriginalDB()
	var stolen reader.DBReader = encoder.NewStolenDB()

	filename := flag.String("f", "", "filename.[xml/json]")
	flag.Parse()

	if len(*filename) == 0 {
		panic("no such file or directory")
	}

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	switch path.Ext(*filename) {
	case ".xml":
		data := original.Reader(file)
		output, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
	case ".json":
		data := stolen.Reader(file)
		output, err := xml.MarshalIndent(data, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
	default:
		file.Close()
		panic("non valid extension")
	}
}
