package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"
)

var (
	example = []byte(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)

	input = []byte(`2264552475
7681287325
3878781441
6868471776
7175255555
7517441253
3513418848
4628736747
1133155762
8816621663`)
)

func adjacent(x, y int, data [][]int) [][]int {
	h, w := len(data), len(data[0])

	var coords [][]int
	for line := -1; line <= 1; line++ {
		for col := -1; col <= 1; col++ {
			x1, y1 := x+line, y+col
			if (x == x1 && y == y1) || x1 < 0 || x1 >= w || y1 < 0 || y1 >= h {
				continue
			}
			coords = append(coords, []int{x1, y1})
		}
	}

	return coords
}

func increaseAdjecent(x, y int, ns [][]int) int {
	var count int
	for _, a := range adjacent(x, y, ns) {
		x, y := a[0], a[1]
		if ns[y][x] == 10 || ns[y][x] == 0 {
			continue
		}
		ns[y][x]++
		count++
	}
	return count
}

func walkFlashed(levels [][]int, counter *int) {
	for y, line := range levels {
		for x, _ := range line {
			if levels[y][x] == 10 {
				levels[y][x] = 0
				*counter++
				if changed := increaseAdjecent(x, y, levels); changed > 0 {
					walkFlashed(levels, counter)
				}
			}
		}
	}
}

func increase(ns [][]int) {
	for y, line := range ns {
		for x, _ := range line {
			ns[y][x]++
		}
	}
}

func run(r io.Reader) int {
	var levels [][]int

	s := bufio.NewScanner(r)
	for s.Scan() {
		var line []int
		for _, v := range s.Text() { // C way to parse str to int
			line = append(line, int(byte(v)-byte('0')))
		}
		levels = append(levels, line)
	}

	var totalFlashes int
	for step := 1; step <= 100; step++ {
		increase(levels)
		walkFlashed(levels, &totalFlashes)
	}
	return totalFlashes
}

func run2(r io.Reader) int {
	var levels [][]int

	s := bufio.NewScanner(r)
	for s.Scan() {
		var line []int
		for _, v := range s.Text() { // C way to parse str to int
			line = append(line, int(byte(v)-byte('0')))
		}
		levels = append(levels, line)
	}

	allFlashed := func() bool {
		for y, line := range levels {
			for x, _ := range line {
				if levels[y][x] != 0 {
					return false
				}
			}
		}
		return true
	}

	var total, na int
	for {
		total++
		increase(levels)
		walkFlashed(levels, &na)
		if allFlashed() {
			break
		}
	}
	return total
}

func print(f [][]int) {
	for y, line := range f {
		for x, _ := range line {
			fmt.Printf("%v ", f[y][x])
		}
		fmt.Println()
	}
	fmt.Println("----------")
}

func TestExample(t *testing.T) {
	expect := 1656

	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 1632

	r := bytes.NewReader(input)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 195

	r := bytes.NewReader(example)
	if got := run2(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 303

	r := bytes.NewReader(input)
	if got := run2(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
