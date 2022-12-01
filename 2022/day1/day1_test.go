package day1

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sort"
	"strconv"
	"testing"
)

var (
	example = []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
)

func sum(ns []int) int {
	var a int
	for _, n := range ns {
		a += n
	}
	return a
}

func insert(a []int, v int) []int {
	a = append(a, v)
	sort.Ints(a)
	if len(a) > 3 {
		a = a[len(a)-3:]
	}
	return a
}

func max(r io.Reader) (max []int) {
	sum := 0
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			max = insert(max, sum)
			sum = 0
			continue
		}
		sum = sum + n
	}
	max = insert(max, sum)
	return max
}

func TestExampleMax(t *testing.T) {
	expect := 24000

	r := bytes.NewReader(example)
	if m := max(r); m[2] != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}

func TestMax(t *testing.T) {
	expect := 66616

	f, _ := os.Open("input")
	if m := max(f); m[2] != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}

func TestExampleTotal(t *testing.T) {
	expect := 45000

	r := bytes.NewReader(example)
	if m := sum(max(r)); m != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}

func TestTotal(t *testing.T) {
	expect := 199172

	f, _ := os.Open("input")
	if m := sum(max(f)); m != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}
