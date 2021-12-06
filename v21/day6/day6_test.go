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
			n, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			ns = append(ns, n)
		}
	}
	return ns
}

func run(r io.Reader, days int) int {

	fishes := make([]int, 9) // track 9 fish at most (indexed 0-8)
	for _, n := range read(r) {
		fishes[n] += 1
	}

	for d := 1; d <= days; d++ {
		tmp := make([]int, len(fishes))

		for i := 1; i <= 8; i++ { // shiftleft
			tmp[i-1] = fishes[i]
		}

		if fishes[0] > 0 {
			tmp[6] += fishes[0] // add fish0 counts to fish6 counts
			tmp[8] = fishes[0]  // amount of new fish created
		}

		fishes = tmp
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
