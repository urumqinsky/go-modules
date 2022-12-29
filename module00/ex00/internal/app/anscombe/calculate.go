package anscombe

import "math"

// Mean Среднее арифметическое — определяется как число, равное сумме всех чисел
// множества, делённой на их количество.
func (a *Anscombe) Mean() float64 {
	if len(a.Data) == 0 {
		return math.NaN()
	}
	return float64(a.Sum) / float64(len(a.Data))
}

// Median Серединное значение упорядоченного набора чисел.
func (a *Anscombe) Median() float64 {
	length := len(a.Data)
	if length == 0 {
		return math.NaN()
	} else if length%2 == 0 {
		return float64(a.Data[length/2-1]+a.Data[length/2+1]) / 2
	} else {
		return float64(a.Data[length/2])
	}
}

// Mode Значение во множестве наблюдений, которое встречается наиболее часто.
func (a *Anscombe) Mode() float64 {
	var mode, max int
	for key, val := range a.ModeNum {
		if val > max || (val == max && key < mode) {
			mode = key
			max = val
		}
	}
	return float64(mode)
}

// SD Среднеквадратическое отклонение — наиболее распространённый показатель
// рассеивания значений случайной величины относительно её математического
// ожидания (аналога среднего арифметического с бесконечным числом исходов).
func (a *Anscombe) SD() float64 {
	sd := 0.00
	mean := a.Mean()
	for i := 0; i < len(a.Data); i++ {
		sd += math.Pow(float64(a.Data[i])-mean, 2)
	}
	return math.Sqrt(sd / float64(len(a.Data)))
}
