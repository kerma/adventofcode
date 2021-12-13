package day13

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	p1      = regexp.MustCompile(`(\d+),(\d+)`)
	p2      = regexp.MustCompile(`fold along (\w)=(\d+)`)
	example = []byte(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`)
)

func run(r io.Reader) int {

	var dots = make(map[string]bool)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if m := p1.FindStringSubmatch(line); len(m) > 0 {
			dots[m[0]] = true
		}
	}

	var folds []string
	var foldAlong []int
	for scanner.Scan() {
		if m := p2.FindStringSubmatch(scanner.Text()); len(m) > 0 {
			folds = append(folds, m[1])
			n, _ := strconv.Atoi(m[2])
			foldAlong = append(foldAlong, n)
		}
	}

	var count int
	for i, xory := range folds {
		for k, _ := range dots {
			x, y := coords(k)
			if xory == "y" {
				if y <= foldAlong[i] {
					continue
				}
				delete(dots, k)

				distance := y - foldAlong[i]
				y2 := foldAlong[i] - distance
				dots[key(x, y2)] = true
			} else if xory == "x" {
				if x <= foldAlong[i] {
					continue
				}
				delete(dots, k)

				distance := x - foldAlong[i]
				x2 := foldAlong[i] - distance
				dots[key(x2, y)] = true
			}
		}
		if i == 0 { // count only after first fold
			count = len(dots)
		}
	}
	Print(dots)
	return count
}

func coords(s string) (int, int) {
	m := strings.Split(s, ",")
	x, _ := strconv.Atoi(m[0])
	y, _ := strconv.Atoi(m[1])
	return x, y
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func Print(dots map[string]bool) {
	var mx, my int
	for k, _ := range dots {
		x, y := coords(k)
		if x >= mx {
			mx = x + 1
		}
		if y >= my {
			my = y + 1
		}
	}
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			if _, ok := dots[key(x, y)]; ok {
				fmt.Printf("%s", "\u2588")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func TestExample(t *testing.T) {
	expect := 17

	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOneAndTwo(t *testing.T) {
	expect := 610

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	if got := run(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
