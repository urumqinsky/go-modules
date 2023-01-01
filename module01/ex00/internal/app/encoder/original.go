package encoder

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/database"
)

type OriginalDB struct{}

func NewOriginalDB() *OriginalDB {
	return &OriginalDB{}
}

func (o *OriginalDB) Reader(file *os.File) database.Recipes {
	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var data = database.Recipes{}
	err = xml.Unmarshal(input, &data)
	if err != nil {
		panic(err)
	}
	return data
}
