package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/exp/slices"

	. "kerma/adventofcode/2022/aoc"
)

var example = []byte(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`)

type Point struct {
	x    int
	y    int
	sand bool
}

func NewPoint(s, sep string) Point {
	p := strings.Split(s, sep)
	x, err := strconv.Atoi(p[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(p[1])
	if err != nil {
		panic(err)
	}
	return Point{x: x, y: y}
}

func SortX(a, b Point) bool {
	return a.x < b.x
}

func SortY(a, b Point) bool {
	return a.y < b.y
}

func (p *Point) Draw() string {
	if p.sand {
		return "o"
	}
	return "#"
}

func Get(ps []Point, x, y int) *Point {
	for _, p := range ps {
		if p.x == x && p.y == y {
			return &p
		}
	}
	return nil
}

func ProduceRocks(path []Point) []Point {
	rocks := make([]Point, len(path))
	copy(rocks, path)

	for {
		if len(path) < 2 {
			break
		}
		a, b := path[0], path[1]

		if a.y == b.y { // points from left to right
			ps := []Point{a, b}
			slices.SortFunc(ps, SortX)
			a, b = ps[0], ps[1]

			dx := b.x - a.x
			for x := 1; x < dx; x++ {
				rocks = append(rocks, Point{x: a.x + x, y: a.y})
			}
		}

		if a.x == b.x { // points from top to bottom
			ps := []Point{a, b}
			slices.SortFunc(ps, SortY)
			a, b = ps[0], ps[1]

			dy := b.y - a.y
			for y := 1; y < dy; y++ {
				rocks = append(rocks, Point{x: a.x, y: a.y + y})
			}

		}

		path = path[1:]
	}

	return rocks
}

func Drip(rocks []Point, max int) []Point {
	var x, y = 500, 0
	for {
		y++
		if y > max {
			break
		}

		p := Get(rocks, x, y)
		if p == nil {
			continue // falling
		}

		// try fall left
		x = x - 1
		p = Get(rocks, x, y)
		if p == nil {
			continue // falling
		}

		// point exists down left, try fall right instead
		x = x + 2
		p = Get(rocks, x, y)
		if p == nil {
			continue // falling
		}

		// cannot fall, stay
		rocks = append(rocks, Point{x: x - 1, y: y - 1, sand: true})
		break
	}
	return rocks
}

func Print(rocks []Point) {
	slices.SortFunc(rocks, SortX)
	sx := rocks[0].x - 1
	_, max := Last(rocks)

	slices.SortFunc(rocks, SortY)
	ry := rocks[len(rocks)-1].y + 1

	fmt.Print("\t")
	for x := sx; x <= max.x; x++ {
		if x == 500 {
			fmt.Print("+")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
	for y := 0; y <= ry; y++ {
		fmt.Printf("%d:\t", y)
		for x := sx; x <= max.x; x++ {
			if p := Get(rocks, x, y); p != nil {
				fmt.Print(p.Draw())
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func run(r io.Reader) int {
	s := bufio.NewScanner(r)

	var rocks []Point
	for s.Scan() {
		line := strings.Split(s.Text(), " -> ")
		var path []Point
		for _, i := range line {
			p := NewPoint(i, ",")
			path = append(path, p)
		}
		rocks = append(rocks, ProduceRocks(path)...)
	}

	slices.SortFunc(rocks, SortY)
	_, max := Last(rocks)

	total := 0
	for {
		rocks = Drip(rocks, max.y)
		sand := Filter(rocks, func(p Point) bool {
			return p.sand
		})
		if len(sand) > total {
			total = len(sand)
			continue
		}
		break
	}

	// Print(rocks)

	return total
}

func TestExample(t *testing.T) {
	expect := 24
	r := bytes.NewReader(example)
	if s := run(r); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestInput(t *testing.T) {
	expect := 1001

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
