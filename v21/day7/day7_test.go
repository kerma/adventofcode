package day7

import (
	"bytes"
	"io"
	"os"
	"sort"
	"testing"

	"github.com/kerma/adventofcode/v21/util"
)

var (
	example = []byte(`16,1,2,0,4,2,7,1,2,14`)
)

func fuelCostTo(ns []int, t int, asc bool) int {
	var c int
	if asc == true {
		for _, n := range ns {
			for ii := 1; n+ii <= t; ii++ {
				c = c + ii
			}
		}
	} else {
		for _, n := range ns {
			for ii := 1; n-ii >= t; ii++ {
				c = c + ii
			}
		}
	}
	return c
}

func run(r io.Reader) int {
	ns := util.ReadInts(r)
	sort.Ints(ns)

	var c int
	m := util.Median(ns)
	for _, n := range ns {
		c += util.Abs(n - m)
	}

	return c
}

func run2(r io.Reader) int {
	ns := util.ReadInts(r)
	sort.Ints(ns)

	m := util.Median(util.Uniq(ns))
	i := util.FindIndex(m, ns)
	a, b := ns[0:i+1], ns[i+1:]

	costs := make([]int, 0)
	t1, t2 := a[len(a)-1], b[0]
	diff := t2 - t1
	for i := 0; i <= diff; i++ {
		n := diff - i
		ca := fuelCostTo(a, t1+i, true)
		cb := fuelCostTo(b, t2-n, false)
		costs = append(costs, ca+cb)
	}

	sort.Ints(costs)
	return costs[0]
}

func TestExample(t *testing.T) {
	expect := 37
	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 168
	r := bytes.NewReader(example)
	if got := run2(r); expect != got {
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

	if got := run(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 95851339

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run2(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
