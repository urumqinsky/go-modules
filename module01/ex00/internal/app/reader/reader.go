package reader

import (
	"os"

	"github.com/urumqinsky/go-modules/module01/ex00/internal/app/database"
)

type DBReader interface {
	Reader(file *os.File) database.Recipes
}
