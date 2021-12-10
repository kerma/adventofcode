package day9

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sort"
	"testing"
)

var (
	example = []byte(`2199943210
3987894921
9856789892
8767896789
9899965678
`)
)

func adjacent(x, y int, data [][]int) []int {
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

func isLow(n int, ps []int) bool {
	for _, p := range ps {
		if n >= p {
			return false
		}
	}
	return true
}

func run(r io.Reader) int {

	var fields [][]int
	s := bufio.NewScanner(r)
	for s.Scan() {
		var line []int
		for _, v := range s.Text() { // C way to parse str to int
			line = append(line, int(byte(v)-byte('0')))
		}
		fields = append(fields, line)
	}

	var sum, count int
	for y, line := range fields {
		for x, n := range line {
			a := adjacent(x, y, fields)
			if isLow(n, a) {
				sum += n
				count++
			}
		}
	}

	return sum + count
}

/// Part two

var ( // globals, so that we can pass less things around in recursive findBasinSize
	fields  [][]int
	visited [][]bool
)

func findBasinSize(x, y int) int {
	if fields[y][x] == 9 || visited[y][x] {
		return 0
	}
	visited[y][x] = true

	size := 1
	if y > 0 {
		size += findBasinSize(x, y-1)
	}
	if y < len(fields)-1 {
		size += findBasinSize(x, y+1)
	}
	if x > 0 {
		size += findBasinSize(x-1, y)
	}
	if x < len(fields[0])-1 {
		size += findBasinSize(x+1, y)
	}

	return size
}

func run2(r io.Reader) int {

	s := bufio.NewScanner(r)
	for s.Scan() {
		var line []int
		for _, v := range s.Text() { // C way to parse str to int
			line = append(line, int(byte(v)-byte('0')))
		}
		fields = append(fields, line)
	}

	visited = make([][]bool, len(fields))
	for i := range visited { // make sure to allocate rightly sized array
		visited[i] = make([]bool, len(fields[0]))
	}

	var sizes []int
	for y, line := range fields {
		for x, n := range line {
			a := adjacent(x, y, fields)
			if isLow(n, a) {
				sizes = append(sizes, findBasinSize(x, y))
			}
		}
	}

	sort.Ints(sizes)
	prod := 1
	for _, x := range sizes[len(sizes)-3:] {
		prod = prod * x
	}
	return prod
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

func TestExamplePartTwo(t *testing.T) {
	expect := 1134
	r := bytes.NewReader(example)

	// reset global vars
	fields = make([][]int, 0)
	visited = make([][]bool, 0)

	if got := run2(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 949905

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// reset global vars
	fields = make([][]int, 0)
	visited = make([][]bool, 0)

	if got := run2(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
