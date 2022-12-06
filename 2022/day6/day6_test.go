package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/exp/slices"
)

func search(x []byte, target byte) (int, bool) {
	for i, c := range x {
		if c == target {
			return i, true
		}
	}
	return -1, false
}

func run(r io.Reader, count int) int {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	var (
		pos   int
		stack []byte
	)
	for i, c := range b {
		if len(stack) == count {
			pos = i
			break
		}

		if i, ok := search(stack, c); ok {
			stack = slices.Delete(stack, 0, i+1)
		}

		stack = append(stack, c)
	}
	return pos
}

func TestExamples(t *testing.T) {
	testCases := []struct {
		data   []byte
		count  int
		expect int
	}{
		{
			data:   []byte(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`),
			count:  4,
			expect: 7,
		},
		{
			data:   []byte(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`),
			count:  14,
			expect: 19,
		},
		{
			data:   []byte(`bvwbjplbgvbhsrlpgdmjqwftvncz`),
			count:  4,
			expect: 5,
		},
		{
			data:   []byte(`bvwbjplbgvbhsrlpgdmjqwftvncz`),
			count:  14,
			expect: 23,
		},
		{
			data:   []byte(`nppdvjthqldpwncqszvftbrmjlhg`),
			count:  4,
			expect: 6,
		},
		{
			data:   []byte(`nppdvjthqldpwncqszvftbrmjlhg`),
			count:  14,
			expect: 23,
		},
		{
			data:   []byte(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`),
			count:  4,
			expect: 10,
		},
		{
			data:   []byte(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`),
			count:  14,
			expect: 29,
		},
		{
			data:   []byte(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`),
			count:  4,
			expect: 11,
		},
		{
			data:   []byte(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`),
			count:  14,
			expect: 26,
		},
	}
	for _, tc := range testCases {
		t.Run(string(tc.data), func(t *testing.T) {
			r := bytes.NewReader(tc.data)
			if s := run(r, tc.count); s != tc.expect {
				t.Fatalf("%d != %d\n", s, tc.expect)
			}
		})
	}
}

func TestStartOfPacket(t *testing.T) {
	expect := 1850

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, 4); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestStartOfMessage(t *testing.T) {
	expect := 2823

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, 14); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
