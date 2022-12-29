package encoder

import (
	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/database"
	"os"
)

type OriginalDB struct{}

func (s *OriginalDB) Reader(file *os.File) database.Recipes {
	return database.Recipes{}
}
