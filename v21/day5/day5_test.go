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

func key(x, y int) string {
	return fmt.Sprintf("%d|%d", x, y)
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

func findPoints(a, b coordinate) []string {
	keys := make([]string, 0)
	if a.x == b.x { // vert
		if a.y <= b.y {
			for i := a.y; i <= b.y; i++ {
				keys = append(keys, key(a.x, i))
			}
		} else {
			for i := b.y; i <= a.y; i++ {
				keys = append(keys, key(a.x, i))
			}
		}
	} else if a.y == b.y { // horiz
		if a.x <= b.x {
			for i := a.x; i <= b.x; i++ {
				keys = append(keys, key(i, a.y))
			}
		} else {
			for i := b.x; i <= a.x; i++ {
				keys = append(keys, key(i, a.y))
			}
		}
	}
	return keys
}

func findDiagonalPoints(a, b coordinate) []string {
	xs := make([]int, 0)
	ys := make([]int, 0)
	i := a.x
	for {
		if i == b.x {
			break
		}
		if i > b.x {
			i--
		} else {
			i++
		}
		xs = append(xs, i)
	}
	i = a.y
	for {
		if i == b.y {
			break
		}
		if i > b.y {
			i--
		} else {
			i++
		}
		ys = append(ys, i)
	}

	keys := make([]string, 0)
	keys = append(keys, key(a.x, a.y))
	for i, x := range xs {
		keys = append(keys, key(x, ys[i]))
	}
	return keys
}

func isDiagonal(a, b coordinate) bool {
	if a.x == b.x || a.y == b.y {
		return false
	}
	return true
}

func findPointsBetween(a, b coordinate, diagonal bool) []string {
	if isDiagonal(a, b) {
		if diagonal {
			return findDiagonalPoints(a, b)
		}
		return []string{}
	}
	return findPoints(a, b)
}

func run(r io.Reader, diagonal bool) int {
	counter := make(map[string]int, 0)

	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		m := pattern.FindStringSubmatch(scanner.Text())
		c1, c2 := parseTo(m)
		keys := findPointsBetween(c1, c2, diagonal)
		for _, k := range keys {
			counter[k] += 1
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
