package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"testing"
)

var (
	pattern = regexp.MustCompile(`([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)`)
	example = []byte(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`)
)

type coordinate struct {
	x, y int
}

func (c coordinate) key() string {
	return fmt.Sprintf("%d|%d", c.x, c.y)
}

func parseTo(ss []string) (coordinate, coordinate) {
	var err error
	r := make([]int, len(ss)-1)
	for i, s := range ss[1:] {
		r[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return coordinate{r[0], r[1]}, coordinate{r[2], r[3]}

}

func findPoints(a, b coordinate) []coordinate {
	cs := make([]coordinate, 0)
	if a.x == b.x { // vert
		r1, r2 := b.y, a.y
		if a.y <= b.y {
			r1, r2 = a.y, b.y
		}
		for i := r1; i <= r2; i++ {
			cs = append(cs, coordinate{a.x, i})
		}
	} else if a.y == b.y { // horiz
		r1, r2 := b.x, a.x
		if a.x <= b.x {
			r1, r2 = a.x, b.x
		}
		for i := r1; i <= r2; i++ {
			cs = append(cs, coordinate{i, a.y})
		}
	}
	return cs
}

func findDiagonalPoints(a, b coordinate) []coordinate {
	x := a.x
	y := a.y
	xs := []int{x}
	ys := []int{y}
	for {
		if x == b.x {
			break
		}
		if x > b.x {
			x--
		} else {
			x++
		}
		xs = append(xs, x)
		if y > b.y {
			y--
		} else {
			y++
		}
		ys = append(ys, y)
	}

	cs := make([]coordinate, 0)
	for i, x := range xs {
		cs = append(cs, coordinate{x, ys[i]})
	}
	return cs
}

func isDiagonal(a, b coordinate) bool {
	if a.x == b.x || a.y == b.y {
		return false
	}
	return true
}

func findPointsBetween(a, b coordinate, diagonal bool) []coordinate {
	if isDiagonal(a, b) {
		if diagonal {
			return findDiagonalPoints(a, b)
		}
		return []coordinate{}
	}
	return findPoints(a, b)
}

func run(r io.Reader, diagonal bool) int {
	counter := make(map[string]int, 0)

	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		m := pattern.FindStringSubmatch(scanner.Text())
		c1, c2 := parseTo(m)
		points := findPointsBetween(c1, c2, diagonal)
		for _, p := range points {
			counter[p.key()] += 1
		}
	}

	var c int
	for _, v := range counter {
		if v >= 2 {
			c++
		}
	}

	return c
}

func TestExample(t *testing.T) {
	expect := 5
	r := bytes.NewReader(example)
	if got := run(r, false); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 12
	r := bytes.NewReader(example)
	if got := run(r, true); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 5167

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file, false); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 17604

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file, true); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
