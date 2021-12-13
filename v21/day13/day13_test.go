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

	// var instructions [][]int
	var dots = make(map[string]bool)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if m := p1.FindStringSubmatch(line); len(m) > 0 {
			dots[m[0]] = true
			// x, _ := strconv.Atoi(m[1])
			// y, _ := strconv.Atoi(m[2])
			// instructions = append(instructions, []int{x, y})
		}
	}
	// fmt.Println(instructions)

	var f1 []string
	var f2 []int
	for scanner.Scan() {
		if m := p2.FindStringSubmatch(scanner.Text()); len(m) > 0 {
			f1 = append(f1, m[1])
			n, _ := strconv.Atoi(m[2])
			f2 = append(f2, n)
		}
	}
	fmt.Println(f1)
	fmt.Println(f2)

	coords := func(s string) (int, int) {
		m := strings.Split(s, ",")
		x, _ := strconv.Atoi(m[0])
		y, _ := strconv.Atoi(m[1])
		return x, y
	}
	key := func(x, y int) string {
		return fmt.Sprintf("%d,%d", x, y)
	}

	foldx, foldy := false, false
	if f1[0] == "x" {
		foldx = true
	}
	if f1[0] == "y" {
		foldy = true
	}
	foldfrom := f2[0]

	for k, _ := range dots {
		x, y := coords(k)
		if foldy {
			if y < foldfrom {
				continue
			}
			delete(dots, k)

			distance := y - foldfrom
			y2 := foldfrom - distance
			dots[key(x, y2)] = true
		}
		if foldx {
			if x < foldfrom {
				continue
			}
			delete(dots, k)

			distance := x - foldfrom
			x2 := foldfrom - distance
			dots[key(x2, y)] = true
		}
	}
	// fmt.Println(dots)
	// fmt.Println(len(dots))

	// for _, line := range lines {
	// }

	return len(dots)
}

func print(coords [][]int) {
	for _, line := range coords {
		for _, n := range line {
			if n == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
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

func TestPartOne(t *testing.T) {
	expect := -1

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
// 	expect := -1
// 	r := bytes.NewReader(example)
// 	if got := run2(r); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }

// func TestPartTwo(t *testing.T) {
// 	expect := 91533

// 	file, err := os.Open("input")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer file.Close()

// 	if got := run2(file); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }
