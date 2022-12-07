package aoc

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
