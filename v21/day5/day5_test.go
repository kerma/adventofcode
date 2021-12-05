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

func filterInt(ss []string) []int {
	var err error
	r := make([]int, len(ss)-1)
	for i, s := range ss {
		if i == 0 {
			continue
		}
		r[i-1], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return r
}

func key(x, y int) string {
	return fmt.Sprintf("%d|%d", x, y)
}

func getPoints(point []int) []string {

	var x1, y1, x2, y2 = point[0], point[1], point[2], point[3]
	keys := make([]string, 0)
	if x1 == x2 { // vert
		if y1 <= y2 {
			for i := y1; i <= y2; i++ {
				keys = append(keys, key(x1, i))
			}
		} else {
			for i := y2; i <= y1; i++ {
				keys = append(keys, key(x1, i))
			}
		}
	} else if y1 == y2 { // horiz
		if x1 <= x2 {
			for i := x1; i <= x2; i++ {
				keys = append(keys, key(i, y1))
			}
		} else {
			for i := x2; i <= x1; i++ {
				keys = append(keys, key(i, y1))
			}
		}
	}
	return keys
}

func getPointsV(point []int) []string {
	var x1, y1, x2, y2 = point[0], point[1], point[2], point[3]
	xs := make([]int, 0)
	ys := make([]int, 0)
	i := x1
	for {
		if i == x2 {
			break
		}
		if i > x2 {
			i--
		} else {
			i++
		}
		xs = append(xs, i)
	}
	i = y1
	for {
		if i == y2 {
			break
		}
		if i > y2 {
			i--
		} else {
			i++
		}
		ys = append(ys, i)
	}
	keys := make([]string, 0)
	keys = append(keys, key(x1, y1))
	for i, x := range xs {
		keys = append(keys, key(x, ys[i]))
	}
	return keys
}

func isDiagonal(point []int) bool {
	var x1, y1, x2, y2 = point[0], point[1], point[2], point[3]
	if x1 == x2 || y1 == y2 {
		return false
	}
	return true
}

func calculateKeys(point []int, diagonal bool) []string {
	if isDiagonal(point) {
		if diagonal {
			return getPointsV(point)
		} else {
			return []string{}
		}
	}
	return getPoints(point)
}

func run(r io.Reader, diagonal bool) int {
	mapCounter := make(map[string]int, 0)
	scanner := bufio.NewScanner(r)

	for i := 0; scanner.Scan(); i++ {
		m := pattern.FindStringSubmatch(scanner.Text())
		point := filterInt(m)
		keys := calculateKeys(point, diagonal)
		for _, k := range keys {
			mapCounter[k] += 1
		}
	}

	var c int
	for _, v := range mapCounter {
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
