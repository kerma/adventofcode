package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	. "kerma/adventofcode/2022/aoc"
)

var example = []byte(`30373
25512
65332
33549
35390`)

func key(x, y int) string {
	return fmt.Sprintf("%d-%d", y, x)
}

func score(grid [][]int) int {
	width := len(grid[0])
	height := len(grid)

	calc := func(v, x, y int) int {
		counter := 0
		d := make([]int, 4)

		for i := x + 1; i < width; i++ {
			counter++
			if v <= grid[y][i] {
				break
			}
		}
		d[0] = counter
		counter = 0

		for i := x - 1; i >= 0; i-- {
			counter++
			if v <= grid[y][i] {
				break
			}
		}
		d[1] = counter
		counter = 0

		for i := y + 1; i < height; i++ {
			counter++
			if v <= grid[i][x] {
				break
			}
		}
		d[2] = counter
		counter = 0

		for i := y - 1; i >= 0; i-- {
			counter++
			if v <= grid[i][x] {
				break
			}
		}
		d[3] = counter

		return Prod(d)
	}

	var res int
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			score := calc(grid[y][x], x, y)
			if score > res {
				res = score
			}
		}
	}

	return res
}

func visible(grid [][]int) int {
	counts := make(map[string]int)

	count := func(left, right []int, key func(int) string) {
		var prev, loc int
		for i, v := range left {
			if i == 0 {
				prev = v
			}
			if v > prev {
				counts[key(loc)] = v
				prev = v
			}
			loc++
		}
		loc = len(right)
		for i, v := range right {
			if i == 0 {
				prev = v
			}
			if v > prev {
				counts[key(loc)] = v
				prev = v
			}
			loc--
		}
	}

	// iterate rows, split row to left and right and iterate both parts
	// look for trees which are taller than the previous one, store those
	// values in counts map.
	// skip the first and last rows, will count them later.
	for y := 1; y < len(grid)-1; y++ {
		left, right := grid[y], grid[y]
		left = left[:len(left)-1]
		right = Reverse(right)
		right = right[:len(right)-1]

		count(left, right, func(a int) string {
			return fmt.Sprintf("%d-%d", y, a)
		})
	}

	// iterate all columns, do the same as for rows
	h := len(grid)
	for x := 1; x < len(grid[0])-1; x++ {
		var col []int
		for k := 0; k < h; k++ {
			col = append(col, grid[k][x])
		}

		left, right := col, col
		right = Reverse(right)
		left = left[:len(left)-1]
		right = right[:len(right)-1]

		count(left, right, func(a int) string {
			return fmt.Sprintf("%d-%d", a, x)
		})
	}

	visible := []int{
		2 * len(grid),          // everything in first and last row
		2 * (len(grid[0]) - 2), // first and last columns without top and bottom,
		len(counts),
	}

	return Sum(visible)
}

func run(r io.Reader, solve func([][]int) int) int {
	s := bufio.NewScanner(r)

	var grid [][]int
	for s.Scan() {
		grid = append(grid, BytesToInts(s.Bytes()))
	}
	return solve(grid)

}

func TestExample(t *testing.T) {
	expect := 21
	r := bytes.NewReader(example)
	if s := run(r, visible); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestExampleScore(t *testing.T) {
	expect := 8
	r := bytes.NewReader(example)
	if s := run(r, score); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestVisible(t *testing.T) {
	expect := 1703

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, visible); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestScore(t *testing.T) {
	expect := 496650

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, score); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
