package encoder

import (
	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/database"
	"os"
)

type StolenDB struct{}

func (s *StolenDB) Reader(file *os.File) database.Recipes {
	return database.Recipes{}
}
