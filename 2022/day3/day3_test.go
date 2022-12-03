package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
)

var (
	example = []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)
)

type sack []byte

var priorities = func() map[byte]int {
	lower := byte('a')
	p := make(map[byte]int, 52)
	for i := 1; i <= 26; i++ {
		p[lower] = i
		lower++
	}
	upper := byte('A')
	for i := 27; i <= 52; i++ {
		p[upper] = i
		upper++
	}
	return p
}()

func sum(ns []int) int {
	var a int
	for _, n := range ns {
		a += n
	}
	return a
}

func findCommons(a sack, ss ...sack) []byte {
	b := ss[0]

	set := make(map[byte]struct{})
	for _, i := range a {
		for _, k := range b {
			if i == k {
				set[i] = struct{}{}
			}
		}
	}

	res := make([]byte, 0, len(set))
	for k := range set {
		res = append(res, k)
	}

	if len(ss) == 1 {
		return res
	}
	return findCommons(res, ss[1:]...)
}

func compartments(r io.Reader) int {
	s := bufio.NewScanner(r)
	commons := make([]byte, 0)
	for s.Scan() {
		b := s.Bytes()
		m := int(len(b) / 2)
		a, b := sack(b[:m]), sack(b[m:])
		commons = append(commons, findCommons(a, b)...)
	}

	res := make([]int, 0)
	for _, c := range commons {
		res = append(res, priorities[c])
	}

	return sum(res)
}

func groups(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	commons := make([]byte, 0)
	tmp := make([]sack, 0)
	for scanner.Scan() {
		if len(tmp) == 3 {
			commons = append(commons, findCommons(tmp[0], tmp[1:]...)...)
			tmp = make([]sack, 0)
		}
		b := []byte(scanner.Text()) // force a copy to be stored in tmp
		tmp = append(tmp, sack(b))
	}
	commons = append(commons, findCommons(tmp[0], tmp[1:]...)...)

	res := make([]int, 0)
	for _, c := range commons {
		res = append(res, priorities[c])
	}

	return sum(res)
}

func TestExample(t *testing.T) {
	expect := 157

	r := bytes.NewReader(example)
	if s := compartments(r); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestCompartments(t *testing.T) {
	expect := 8153

	f, _ := os.Open("input")
	defer f.Close()
	if s := compartments(f); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestGroupsExample(t *testing.T) {
	expect := 70

	r := bytes.NewReader(example)
	if s := groups(r); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestGroups(t *testing.T) {
	expect := 2342

	f, _ := os.Open("input")
	defer f.Close()
	if s := groups(f); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
