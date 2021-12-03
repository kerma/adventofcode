package day3

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"testing"
)

var (
	bit0    = byte(48)
	bit1    = byte(49)
	example = []byte(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)
)

func calculate(r io.Reader, size int) int {
	gamma := make([]int, size)
	epsilon := make([]int, size)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for pos, char := range scanner.Bytes() {
			if char == bit1 {
				gamma[pos]++
			} else if char == bit0 {
				epsilon[pos]++
			} else {
				panic("unexpected char")
			}
		}
	}
	gr := ""
	er := ""
	for i, g := range gamma {
		if int(g) == int(epsilon[i]) {
			panic("unexpected counts")
		}
		if int(g) > int(epsilon[i]) {
			gr = gr + "1"
			er = er + "0"
		} else {
			gr = gr + "0"
			er = er + "1"
		}
	}

	grate, err := strconv.ParseInt(gr, 2, 64)
	if err != nil {
		panic(err)
	}
	erate, err := strconv.ParseInt(er, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(grate * erate)
}

func TestExample(t *testing.T) {
	expect := 198
	r := bytes.NewReader(example)
	if got := calculate(r, 5); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestCalculate(t *testing.T) {
	expect := 3277364

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := calculate(file, 12); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
