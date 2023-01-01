package encoder

import (
	"encoding/json"
	"io"
	"os"

	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/database"
)

type StolenDB struct{}

func NewStolenDB() *StolenDB {
	return &StolenDB{}
}

func (s *StolenDB) Reader(file *os.File) database.Recipes {
	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var data database.Recipes
	err = json.Unmarshal(input, &data)
	if err != nil {
		panic(err)
	}
	return data
}
