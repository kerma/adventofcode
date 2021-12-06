package day6

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	example = []byte(`3,4,3,1,2`)
)

func read(r io.Reader) []int {
	scanner := bufio.NewScanner(r)
	ns := make([]int, 0)
	for scanner.Scan() {
		for _, x := range strings.Split(scanner.Text(), ",") {
			x, err := strconv.Atoi(x)
			if err != nil {
				panic(x)
			}
			ns = append(ns, x)
		}
	}
	return ns
}

func run(r io.Reader, days int) int {
	ns := read(r)
	fishes := make([]int, 9)

	for _, n := range ns {
		fishes[n] += 1
	}

	for i := 1; i <= days; i++ {

		newmap := make([]int, len(fishes))

		for i := 1; i <= 8; i++ { // shiftleft
			newmap[i-1] = fishes[i]
		}

		if fishes[0] > 0 {
			newmap[6] += fishes[0]
			newmap[8] = fishes[0]
		}
		fishes = newmap
	}

	var sum int
	for _, x := range fishes {
		sum += x
	}

	return sum
}

func TestExample(t *testing.T) {
	expect := 26
	r := bytes.NewReader(example)
	if got := run(r, 18); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExample2(t *testing.T) {
	expect := 5934
	r := bytes.NewReader(example)
	if got := run(r, 80); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 26984457539
	r := bytes.NewReader(example)
	if got := run(r, 256); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 363101

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file, 80); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 1644286074024

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file, 256); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
