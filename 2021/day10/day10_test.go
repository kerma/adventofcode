package day10

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sort"
	"testing"

	"github.com/kerma/adventofcode/v21/util"
)

var (
	example = []byte(`
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`)
)

func split(r io.Reader) ([]string, *Stack) {
	incomplete := make([]string, 0)
	corrupted := make(Stack, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		isCorrupted := false
		stack := NewStack()
		for _, item := range Stack(line) {
			if err := stack.Push(item); err != nil {
				corrupted = append(corrupted, item)
				isCorrupted = true
				break
			}
		}
		if !isCorrupted {
			incomplete = append(incomplete, line)
		}
	}
	return incomplete, &corrupted
}

func run(r io.Reader) int {

	scoreMap := map[Item]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	scores := make([]int, 0)

	_, corrupted := split(r)
	for _, item := range *corrupted {
		scores = append(scores, scoreMap[item])
	}

	return util.Sum(scores)
}

func run2(r io.Reader) int {

	scoreMap := map[Item]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	score := func(completion Stack) int {
		var total int
		for _, item := range completion {
			total = total * 5
			total += scoreMap[item]
		}
		return total
	}

	complete := func(line string) Stack {
		reversed := NewStackFrom(line).Reverse()

		comps := make([]Item, 0)
		stack := NewStack()
		for _, item := range *reversed {
			if err := stack.PushReverse(item); err != nil {
				comps = append(comps, pairMap[err.found])
			}
		}

		return comps
	}

	comps := make([]Stack, 0)
	incomplete, _ := split(r)
	for _, line := range incomplete {
		comps = append(comps, complete(line))
	}

	scores := make([]int, len(comps))
	for i, c := range comps {
		scores[i] = score(c)
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

func TestExample(t *testing.T) {
	expect := 26397
	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 367059

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePartTwo(t *testing.T) {
	expect := 288957
	r := bytes.NewReader(example)
	if got := run2(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 1952146692

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run2(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
