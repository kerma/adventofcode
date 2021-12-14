package day14

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var (
	pattern = regexp.MustCompile(`(\w{2}) -> (\w)`)
	example = []byte(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`)
)

func run(r io.Reader, steps int) int {

	var template string
	var rules = make(map[string]string)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if m := pattern.FindStringSubmatch(line); len(m) > 0 {
			rules[m[1]] = m[2]
		} else {
			template = line
		}
	}
	// fmt.Println(template)
	// fmt.Println(rules)

	polymer := template
	for step := 1; step <= steps; step++ {
		var builder strings.Builder
		prev := rune(0)
		for _, r := range polymer {
			if prev == 0 {
				prev = r
				builder.WriteRune(r)
				continue
			}
			if rule, ok := rules[string(prev)+string(r)]; ok {
				builder.WriteString(rule)
				builder.WriteRune(r)
				prev = r
			}
		}
		polymer = builder.String()
	}

	var counts []int
	for _, r := range polymer {
		counts = append(counts, strings.Count(polymer, string(r)))
	}
	sort.Ints(counts)

	return counts[len(counts)-1] - counts[0]
}

func TestExample(t *testing.T) {
	expect := 1588

	r := bytes.NewReader(example)
	if got := run(r, 10); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 2233

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	if got := run(file, 10); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

// func TestPartTwo(t *testing.T) {
// 	expect := 2233

// 	file, err := os.Open("input")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer file.Close()
// 	if got := run(file, 40); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }
