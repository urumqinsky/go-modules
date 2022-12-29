package flag

import (
	"flag"
)

type Flags struct {
	Mean   *bool
	Median *bool
	Mode   *bool
	SD     *bool
}

func (f *Flags) Parse() {
	flag.Parse()
}

func NewFlags() *Flags {
	return &Flags{
		flag.Bool("mean", false, "Среднее арифметическое"),
		flag.Bool("median", false, "Серединное значение"),
		flag.Bool("mode", false, "Значение во множестве наблюдений"),
		flag.Bool("sd", false, "Среднеквадратическое отклонение"),
	}
}
