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

type Knot struct {
	x int
	y int
}

func (h *Knot) Move(d string) {
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

func (t *Knot) Follow(h *Knot) {
	dx := h.x - t.x
	cx := 1
	if dx < 0 {
		cx = -1
	}
	dx = Abs(dx)

	dy := h.y - t.y
	cy := 1
	if dy < 0 {
		cy = -1
	}
	dy = Abs(dy)

	if (dy == 0 && dx == 0) || (dy == 1 && dx == 1) {
		return // same or diagonal
	}

	// h
	// .
	// t
	if dy == 2 && dx == 0 {
		t.y = t.y + cy
	}

	// . . h
	// . . .
	// t . .
	if dy == 2 && dx > 0 {
		t.x = t.x + cx
		t.y = t.y + cy
	}

	// . . h
	// t . .
	if dy == 1 && dx == 2 {
		t.x = t.x + cx
		t.y = t.y + cy
	}

	// t . h
	if dy == 0 && dx == 2 {
		t.x = t.x + cx
	}
}

func (t *Knot) String() string {
	return fmt.Sprintf("%d|%d", t.y, t.x)
}

func follow(kx []*Knot) *Knot {
	if len(kx) == 1 {
		return kx[0]
	}
	head, tail := kx[0], kx[1]
	tail.Follow(head)
	return follow(kx[1:])
}

func run(r io.Reader, count int) int {
	s := bufio.NewScanner(r)

	knots := make([]*Knot, count)
	for i := 0; i < count; i++ {
		knots[i] = &Knot{}
	}

	visited := make(map[string]struct{})
	for s.Scan() {
		var (
			count     int
			direction string
		)
		fmt.Sscanf(s.Text(), "%s %d", &direction, &count)

		for i := 0; i < count; i++ {
			knots[0].Move(direction)
			tail := follow(knots)
			visited[tail.String()] = struct{}{}
		}
	}

	// printExample(knots, 15, -6, -11, 15)

	return len(visited)
}

func TestExample(t *testing.T) {
	expect := 13
	r := bytes.NewReader(example)
	if s := run(r, 2); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestRope(t *testing.T) {
	expect := 6271

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, 2); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestExampleKnots(t *testing.T) {
	expect := 36
	r := bytes.NewReader(example2)
	if s := run(r, 10); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestKnots(t *testing.T) {
	expect := 2458
	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, 10); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func printExample(points []*Knot, y1, y2, x1, x2 int) {
	for y := y1; y > y2; y-- {
		for x := x1; x < x2; x++ {
			nr := "."
			if y == 0 && x == 0 {
				nr = "S"
			}
			for i, item := range points {
				if y == item.y && x == item.x {
					nr = fmt.Sprint(i)
				}
			}
			fmt.Print(nr)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("=================================")
	fmt.Println()
}
