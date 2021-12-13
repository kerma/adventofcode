package day1

import (
	"testing"

	"github.com/kerma/adventofcode/v21/util"
)

var (
	example = []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
)

// measure returns the count of cases where the n-th element in the list
// is bigger than the previous one.
// can also be used for second part of the task as
// a+b+c < b+c+d is the same as a < d
func measure(inp []int, n int) (m int) {
	for i, _ := range inp {
		if i == 0 {
			continue
		}
		if i == len(inp)-n {
			break
		}
		if inp[i-1] < inp[i+n] {
			m++
		}
	}
	return
}

func TestExample(t *testing.T) {
	expect := 7
	if m := measure(example, 0); m != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}

func TestExampleSums(t *testing.T) {
	expect := 5
	if m := measure(example, 2); m != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}

func TestMeasure(t *testing.T) {
	expect := 1529
	ints, err := util.ReadInts("input")
	if err != nil {
		t.Fatal(err)
	}

	if m := measure(ints, 0); m != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}

func TestMeasureSums(t *testing.T) {
	expect := 1567
	ints, err := util.ReadInts("input")
	if err != nil {
		t.Fatal(err)
	}

	if m := measure(ints, 2); m != expect {
		t.Fatalf("%d != %d\n", m, expect)
	}
}
