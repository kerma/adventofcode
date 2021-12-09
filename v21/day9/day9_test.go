package day9

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/kerma/adventofcode/v21/util"
)

var (
	example = []byte(`
2199943210
3987894921
9856789892
8767896789
9899965678
`)
)

func findLows(ns [][]int) []int {

	adjacent := func(x, y int, data [][]int) []int {
		h := len(data)
		w := len(data[0])
		a := make([]int, 0)

		prev := x - 1
		if prev >= 0 {
			a = append(a, data[y][prev])
		}

		next := x + 1
		if next < w {
			a = append(a, data[y][next])
		}

		up := y - 1
		if up >= 0 {
			a = append(a, data[up][x])
		}

		down := y + 1
		if down < h {
			a = append(a, data[down][x])
		}

		return a
	}

	isLow := func(n int, ns []int) bool {
		for _, v := range ns {
			if n >= v {
				return false
			}
		}
		return true
	}

	lows := make([]int, 0)
	for y, line := range ns {
		for x, n := range line {
			a := adjacent(x, y, ns)
			if isLow(n, a) {
				lows = append(lows, n)
			}
		}
	}
	return lows
}

func run(r io.Reader) int {
	ns := util.ReadCharsToInts(r)
	lows := findLows(ns)
	return util.Sum(lows) + len(lows)
}

func TestExample(t *testing.T) {
	expect := 15
	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 518

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
