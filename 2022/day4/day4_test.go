package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"golang.org/x/exp/slices"
)

var (
	example = []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`)
)

func parse(s string) ([]int, []int) {
	var a, b [2]int
	if _, err := fmt.Sscanf(s, "%d-%d,%d-%d", &a[0], &a[1], &b[0], &b[1]); err != nil {
		panic(err)
	}

	p := make([]int, 0)
	for i := a[0]; i <= a[1]; i++ {
		p = append(p, i)
	}
	p2 := make([]int, 0)
	for i := b[0]; i <= b[1]; i++ {
		p2 = append(p2, i)
	}

	return p, p2
}

func order(a, b []int) ([]int, []int) {
	if len(a) > len(b) {
		return b, a
	}
	return a, b
}

func overlaps(a, b []int) int {
	var has int
	for _, v := range a {
		if _, ok := slices.BinarySearch(b, v); ok {
			has++
		}
	}
	if len(a) == has {
		return 1
	}
	return 0
}

func contains(a, b []int) int {
	for _, v := range a {
		if _, ok := slices.BinarySearch(b, v); ok {
			return 1
		}
	}
	return 0
}

func run(r io.Reader, find func([]int, []int) int) int {
	scanner := bufio.NewScanner(r)

	var match int
	for scanner.Scan() {
		a, b := parse(scanner.Text())
		match = match + find(order(a, b))
	}
	return match
}

func TestExample(t *testing.T) {
	expect := 2

	r := bytes.NewReader(example)
	if s := run(r, overlaps); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestOverlaps(t *testing.T) {
	expect := 464

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, overlaps); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestContainsExample(t *testing.T) {
	expect := 4

	r := bytes.NewReader(example)
	if s := run(r, contains); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestContains(t *testing.T) {
	expect := 770

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, contains); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
