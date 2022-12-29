package reader

import (
	"os"
)

type DBReader interface {
	Reader(file *os.File) Recipes
}
