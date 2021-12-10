package day9

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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

type position struct {
	x, y, v int
}

func (p position) String() string { return fmt.Sprintf("[v:%d x:%d y:%d]", p.v, p.x, p.y) }

func adjacent(x, y int, data [][]int) []position {
	h := len(data)
	w := len(data[0])
	a := make(positions, 0)

	prev := x - 1
	if prev >= 0 {
		a = append(a, position{prev, y, data[y][prev]})
	}

	next := x + 1
	if next < w {
		a = append(a, position{next, y, data[y][next]})
	}

	up := y - 1
	if up >= 0 {
		a = append(a, position{x, up, data[up][x]})
	}

	down := y + 1
	if down < h {
		a = append(a, position{x, down, data[down][x]})
	}

	return a
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

	isLow := func(n int, ps []position) bool {
		for _, p := range ps {
			if n >= p.v {
				return false
			}
		}
		return true
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

// func TestExamplePartTwo(t *testing.T) {
// 	expect := 1134
// 	r := bytes.NewReader(example)
// 	if got := run2(r); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }
