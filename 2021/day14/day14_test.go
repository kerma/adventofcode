package day14

import (
	"bufio"
	"bytes"
	"fmt"
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

type Node struct {
	value rune
	next  *Node
	prev  *Node
	count int
}

func (n *Node) Key() string {
	if n.prev == nil {
		return "missing"
	}
	return fmt.Sprintf("%c%c", n.prev.value, n.value)
}

func (n *Node) String() string {
	prev := ""
	if n.prev != nil {
		prev = fmt.Sprintf("%c <- ", n.prev.value)
	}
	next := ""
	if n.next != nil {
		next = fmt.Sprintf(" -> %c", n.next.value)
	}
	return fmt.Sprintf("[%s%c%d%s]", prev, n.value, n.count, next)
}

func get(ns []*Node, r rune) *Node {
	for _, x := range ns {
		if x.value == r {
			return x
		}
	}
	return nil
}

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
	list := make([]*Node, 0)
	var prev *Node
	for i, r := range template {
		if i == 0 {
			prev = &Node{value: r, count: 1}
			list = append(list, prev)
			continue
		}
		if node := get(list, r); node != nil {
			node.prev = node
			prev = node
		} else {
			cur := &Node{value: r, count: 1}
			prev.next = cur
			cur.prev = prev
			list = append(list, cur)
			prev = cur
		}
	}
	fmt.Printf("%s\n", rules)
	fmt.Printf("%s\n", list)

	fmt.Printf("%s\n", template)
	for step := 1; step <= 1; step++ {
		polymer := make([]*Node, 0)
		// copy(polymer, list)
		polymer = append(polymer, list[0])

		for _, node := range list {
			fmt.Println("checking node", node)
			if rule, ok := rules[node.Key()]; ok {
				fmt.Println("processing rule", node.Key(), rule)
				value := []rune(rule)[0]
				count := 0
				if existing := get(list, value); existing != nil {
					count = existing.count
					existing.count += 1
				}
				node.prev.count += 1
				node.count += 1
				n := &Node{
					value: value,
					prev:  node.prev,
					next:  &Node{value: node.value, count: node.count},
					count: 1 + count,
				}
				polymer = append(polymer, n)
			}
		}
		fmt.Println(polymer)
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
// 	expect := 2188189693529

// 	file, err := os.Open("input")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer file.Close()
// 	if got := run(file, 40); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }
