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

var example = []byte(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

var example2 = []byte(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)

type Point struct {
	x       int
	y       int
	nr      int
	visited map[string]struct{}
}

func (h *Point) Move(d string) {
	switch d {
	case "U":
		h.y++
	case "D":
		h.y--
	case "L":
		h.x--
	case "R":
		h.x++
	}
}

func (h *Point) Prev(d string) (int, int) {
	switch d {
	case "U":
		return h.x, h.y - 1
	case "D":
		return h.x, h.y + 1
	case "L":
		return h.x + 1, h.y
	case "R":
		return h.x - 1, h.y
	}
	panic("wtf")
}

func (t *Point) Follow(d string, h *Point) {
	dx := Abs(h.x-t.x) > 1
	dy := Abs(h.y-t.y) > 1
	x, y := h.Prev(d)

	if dy {
		t.x = x
		t.y = y
	}
	if dx {
		t.x = x
		t.y = y
	}
}

func (t *Point) String() string {
	return fmt.Sprintf("%d|%d", t.y, t.x)
}

func run(r io.Reader) int {
	s := bufio.NewScanner(r)

	head := &Point{}
	tail := &Point{}
	visited := make(map[string]struct{})
	for s.Scan() {
		var (
			count     int
			direction string
		)
		fmt.Sscanf(s.Text(), "%s %d", &direction, &count)

		for i := 0; i < count; i++ {
			head.Move(direction)
			tail.Follow(direction, head)
			visited[tail.String()] = struct{}{}
		}
	}

	return len(visited)
}

func TestExample(t *testing.T) {
	expect := 13
	r := bytes.NewReader(example)
	if s := run(r); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestVisit(t *testing.T) {
	expect := 6271

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

// func TestScore(t *testing.T) {
// 	expect := 496650

// 	f, _ := os.Open("input")
// 	defer f.Close()
// 	if s := run(f, score); s != expect {
// 		t.Fatalf("%d != %d\n", s, expect)
// 	}
// }

// func follow(xs []*Head) {
// 	if len(xs) < 2 {
// 		return
// 	}
// 	head, tail := xs[0], xs[1]
// 	tail.Follow(head)
// 	tail.visited[tail.String()] = struct{}{}
// 	follow(xs[1:])
// }

// func run2(r io.Reader) int {
// 	s := bufio.NewScanner(r)

// 	heads := make([]*Head, 10)
// 	for i := 0; i < 10; i++ {
// 		heads[i] = &Head{
// 			nr:      i,
// 			visited: make(map[string]struct{}),
// 		}
// 	}
// 	// head := &Head{}

// 	// tail := &Head{}
// 	head := heads[0]
// 	_, tail := Last(heads)
// 	for s.Scan() {
// 		var (
// 			count     int
// 			direction string
// 		)
// 		fmt.Sscanf(s.Text(), "%s %d", &direction, &count)

// 		fmt.Println(direction)
// 		fmt.Println(count)
// 		for c := 0; c < count; c++ {
// 			head.Move(direction)
// 			follow(heads)
// 		}

// 		for y := 10; y > -10; y-- {
// 			for x := -20; x < 20; x++ {
// 				nr := "."
// 				if y == 0 && x == 0 {
// 					nr = "S"
// 				}
// 				for _, item := range heads {
// 					if y == item.y && x == item.x {
// 						nr = fmt.Sprint(item.nr)
// 					}
// 				}
// 				fmt.Print(nr)
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println("=================================")
// 		fmt.Println()
// 		fmt.Println()
// 		break

// 	}

// 	return len(tail.visited)
// }
