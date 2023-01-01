package anscombe

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/urumqinsky/go-modules/module00/ex00/internal/app/flag"
)

type Anscombe struct {
	Data    []int
	Sum     int
	ModeNum map[int]int
}

func NewAnscombe() *Anscombe {
	return &Anscombe{[]int{}, 0, make(map[int]int)}
}

func (a *Anscombe) Add(num int) {
	a.Data = append(a.Data, num)
}

func (a *Anscombe) AddMean(num int) {
	a.ModeNum[num] += 1
}

func (a *Anscombe) AddSum(num int) {
	a.Sum += num
}

func (a *Anscombe) Scan(scan *bufio.Scanner) {
	for scan.Scan() {
		num, err := strconv.Atoi(scan.Text())
		if err != nil {
			fmt.Println(err)
		} else if num > 100000 || num < -100000 {
			fmt.Println("Value out of range")
		} else {
			a.Add(num)
			a.AddSum(num)
			a.AddMean(num)
		}
	}
	sort.Ints(a.Data)
}

func (a *Anscombe) Print(flags *flag.Flags) {
	if *flags.Mean || len(os.Args) == 1 {
		fmt.Printf("Mean: %f\n", a.Mean())
	}
	if *flags.Median || len(os.Args) == 1 {
		fmt.Printf("Median: %f\n", a.Median())
	}
	if *flags.Mode || len(os.Args) == 1 {
		fmt.Printf("Mode: %f\n", a.Mode())
	}
	if *flags.SD || len(os.Args) == 1 {
		fmt.Printf("SD: %f\n", a.SD())
	}
}
