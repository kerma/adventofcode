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

type grid map[string]int

func key(x, y int) string {
	return fmt.Sprintf("%d|%d", x, y)
}

func run(r io.Reader) int {
	var data [][]int
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		m := pattern.FindStringSubmatch(scanner.Text())
		data = append(data, filterInt(m))
	}

	g := make(grid, 0)
	for _, point := range data {
		var x1, y1, x2, y2 = point[0], point[1], point[2], point[3]
		if x1 == x2 { // vert
			if y1 <= y2 {
				for i := y1; i <= y2; i++ {
					k := key(x1, i)
					g[k] += 1
				}
			} else {
				for i := y2; i <= y1; i++ {
					k := key(x1, i)
					g[k] += 1
				}
			}
		} else if y1 == y2 { // horiz
			if x1 <= x2 {
				for i := x1; i <= x2; i++ {
					k := key(i, y1)
					g[k] += 1
				}
			} else {
				for i := x2; i <= x1; i++ {
					k := key(i, y1)
					g[k] += 1
				}
			}
		}

	}
	var c int
	for _, v := range g {
		if v >= 2 {
			c++
		}
	}

	return c
}

func TestExample(t *testing.T) {
	expect := 5
	r := bytes.NewReader(example)
	if got := run(r); expect != got {
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

	if got := run(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
