package aoc

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sum(ns []int) int {
	var a int
	for _, n := range ns {
		a += n
	}
	return a
}

func SumBy[T any](x []T, pick func(T) int) int {
	var a int
	for _, n := range x {
		a += pick(n)
	}
	return a
}

func Last[T any](x []T) (int, T) {
	i := len(x) - 1
	return i, x[i]
}

func Filter[T any](x []T, filter func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range x {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func Reverse[T any](x []T) []T {
	l := len(x)
	res := make([]T, len(x))
	for i := 0; i < l; i++ {
		res[i] = x[l-i-1]
	}
	return res
}

func Prod(x []int) int {
	var a int
	for i, n := range x {
		if i == 0 {
			a = n
			continue
		}
		a = a * n
	}
	return a
}

func BytesToInts(bs []byte) []int {
	var s []int
	for _, b := range bs { // C way to parse str to int
		s = append(s, int(b-byte('0')))
	}
	return s
}
