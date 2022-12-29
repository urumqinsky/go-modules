package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/database"
	"io"
	"os"
)

// Аннотирование структур

func main() {
	filename := flag.String("f", "", "filename.[xml/json]")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var xmldata database.Recipes
	err = xml.Unmarshal(data, &xmldata)
	if err != nil {
		panic(err)
	}

	data, err = xml.MarshalIndent(xmldata, "   ", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
