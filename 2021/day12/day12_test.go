package day12

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

var (
	pattern = regexp.MustCompile(`(\w+)-(\w+)`)
	example = []byte(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`)

	example2 = []byte(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`)

	input = []byte(`RT-start
bp-sq
em-bp
end-em
to-MW
to-VK
RT-bp
start-MW
to-hr
sq-AR
RT-hr
bp-to
hr-VK
st-VK
sq-end
MW-sq
to-RT
em-er
bp-hr
MW-em
st-bp
to-start
em-st
st-end
VK-sq
hr-st`)
)

type Graph map[string][]string

func run(r io.Reader, twice bool) int {
	g := Graph{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if m := pattern.FindStringSubmatch(s.Text()); len(m) > 0 {
			g[m[1]] = append(g[m[1]], m[2])
			g[m[2]] = append(g[m[2]], m[1])
		}
	}

	var total int
	var visited = []string{}
	walk(&g, "start", visited, twice, &total)

	return total
}

func walk(g *Graph, node string, visited []string, twice bool, counter *int) {

	if node == "end" {
		*counter++
		return
	}

	if contains(visited, node) && isSmall(node) {
		if twice || node == "start" || node == "end" {
			return
		} else {
			twice = true
		}
	}

	visited = append(visited, node)
	for _, next := range (*g)[node] {
		walk(g, next, visited, twice, counter)
	}

}

func isSmall(s string) bool {
	return s == strings.ToLower(s)
}

func contains(hs []string, n string) bool {
	for _, x := range hs {
		if x == n {
			return true
		}
	}
	return false
}

func TestExample(t *testing.T) {
	expect := 10

	r := bytes.NewReader(example)
	if got := run(r, true); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}

	expect = 226
	r = bytes.NewReader(example2)
	if got := run(r, true); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 3463

	r := bytes.NewReader(input)
	if got := run(r, true); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 3509
	r := bytes.NewReader(example2)
	if got := run(r, false); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 91533

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file, false); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
