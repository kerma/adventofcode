package day10

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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

var (
	open = []Item("([{<")
)

type Item rune

func (i Item) IsOpen() bool {
	for _, x := range open {
		if i == x {
			return true
		}
	}
	return false
}

type Stack struct {
	items []Item
}

func NewStack() *Stack {
	return &Stack{
		items: make([]Item, 0),
	}
}

func (s *Stack) Pop() error {
	l := len(s.items)
	if l == 0 {
		return fmt.Errorf("stack empty")
	}
	s.items = s.items[0 : l-1]
	return nil
}

func (s *Stack) Push(item Item) error {
	l := len(s.items)
	if l == 0 {
		if !item.IsOpen() {
			return fmt.Errorf("expected %c, found %c", open, item)
		}
		s.items = append(s.items, item)
		return nil
	}

	last := s.items[l-1]
	switch item {
	case ')':
		if last != '(' {
			return fmt.Errorf("expected %c, found )", last)
		}
		s.Pop()
	case ']':
		if last != '[' {
			return fmt.Errorf("expected %c, found ]", last)
		}
		s.Pop()
	case '}':
		if last != '{' {
			return fmt.Errorf("expected %c, found }", last)
		}
		s.Pop()
	case '>':
		if last != '<' {
			return fmt.Errorf("expected %c, found >", last)
		}
		s.Pop()
	default:
		s.items = append(s.items, item)
	}
	return nil
}

func run(r io.Reader) int {

	scoreMap := map[Item]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	scores := make([]int, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		stack := NewStack()
		for _, r := range []Item(scanner.Text()) {
			if err := stack.Push(r); err != nil {
				scores = append(scores, scoreMap[r])
				break
			}
		}
	}

	return util.Sum(scores)
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

// func TestPartTwo(t *testing.T) {
// 	expect := 95851339

// 	file, err := os.Open("input")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer file.Close()

// 	if got := run2(file); expect != got {
// 		t.Fatalf("%d != %d\n", expect, got)
// 	}
// }
