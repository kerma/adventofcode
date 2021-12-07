package day7

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/kerma/adventofcode/v21/util"
)

var (
	example = []byte(`16,1,2,0,4,2,7,1,2,14`)
)

func run(r io.Reader, two bool) int {

	ns := util.ReadInts(r)
	m := util.Median(ns)

	var c int
	for _, n := range ns {
		c += util.Abs(n - m)
	}

	return c
}

func TestExample(t *testing.T) {
	expect := 37
	r := bytes.NewReader(example)
	if got := run(r, false); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 168
	r := bytes.NewReader(example)
	if got := run(r, true); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 335271

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file, false); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

// func TestPartTwo(t *testing.T) {
// 	expect := 0

// 	file, err := os.Open("input")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer file.Close()

// 	if got := run(file, true); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }
