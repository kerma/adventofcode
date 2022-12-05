package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"

	"golang.org/x/exp/slices"
)

var (
	example = []byte(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)
)

func parseMoves(s string) [3]int {
	var a [3]int
	if _, err := fmt.Sscanf(s, "move %d from %d to %d", &a[0], &a[1], &a[2]); err != nil {
		panic(err)
	}
	return a
}

func parseCrates(ss []string) [][]rune {
	indexes := regexp.MustCompile(`(\d)`).FindAllStringIndex(ss[0], -1)

	var rows [][]rune
	for _, s := range ss[1:] {
		row := make([]rune, len(indexes))
		rs := []rune(s)
		for i, loc := range indexes {
			row[i] = rs[loc[0]]
		}
		rows = append(rows, row)
	}

	var crates [][]rune
	for i := 0; i < len(rows[0]); i++ {
		stack := make([]rune, 0)
		for _, row := range rows {
			if row[i] != ' ' {
				stack = append(stack, row[i])
			}
		}
		crates = append(crates, stack)
	}

	return crates
}

func crateMover(crates [][]rune, moves [3]int) [][]rune {
	count, from, to := moves[0], moves[1]-1, moves[2]-1

	pop := func() rune {
		size := len(crates[from])
		last := crates[from][size-1]
		crates[from] = slices.Delete(crates[from], size-1, size)
		return last
	}

	for i := 0; i < count; i++ {
		crate := pop()
		crates[to] = append(crates[to], crate)
	}

	return crates
}

func crateMover9001(crates [][]rune, moves [3]int) [][]rune {
	count, from, to := moves[0], moves[1]-1, moves[2]-1

	pop := func() []rune {
		size := len(crates[from])
		idx := size - count
		split := crates[from][idx:]
		crates[from] = slices.Delete(crates[from], idx, size)
		return split
	}

	tmp := pop()
	crates[to] = append(crates[to], tmp...)

	return crates
}

func run(r io.Reader, move func([][]rune, [3]int) [][]rune) string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	slices.SortFunc(lines, func(_ string, _ string) bool {
		return true // reverse
	})

	crates := parseCrates(lines)
	for scanner.Scan() {
		moves := parseMoves(scanner.Text())
		crates = move(crates, moves)
	}

	ts := make([]rune, 0)
	for _, s := range crates {
		ts = append(ts, s[len(s)-1])
	}

	return string(ts)
}

func TestExample(t *testing.T) {
	expect := "CMZ"

	r := bytes.NewReader(example)
	if s := run(r, crateMover); s != expect {
		t.Fatalf("%s != %s\n", s, expect)
	}
}

func TestExampleCrateMover9001(t *testing.T) {
	expect := "MCD"

	r := bytes.NewReader(example)
	if s := run(r, crateMover9001); s != expect {
		t.Fatalf("%s != %s\n", s, expect)
	}
}

func TestCrateMover(t *testing.T) {
	expect := "QMBMJDFTD"

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, crateMover); s != expect {
		t.Fatalf("%s != %s\n", s, expect)
	}
}

func TestCrateMover9001(t *testing.T) {
	expect := "NBTVTJNFJ"

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, crateMover9001); s != expect {
		t.Fatalf("%s != %s\n", s, expect)
	}
}
