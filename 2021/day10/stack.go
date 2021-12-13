package day10

import "fmt"

var (
	open    = []Item("([{<")
	pairMap = map[Item]Item{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
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

type ItemErr struct {
	expect Item
	found  Item
}

func NewErr(e, f Item) *ItemErr {
	return &ItemErr{
		expect: e,
		found:  f,
	}
}

func (e *ItemErr) Error() string {
	return fmt.Sprintf("expected %c, found %c", e.expect, e.found)
}

type Stack []Item

func (s *Stack) String() string {
	str := ""
	for _, item := range *s {
		str += fmt.Sprintf("%c", item)
	}
	return str
}

func NewStack() *Stack {
	s := make(Stack, 0)
	return &s
}

func NewStackFrom(str string) *Stack {
	s := Stack(str)
	return &s
}

func (s *Stack) Pop() error {
	l := len(*s)
	if l == 0 {
		return fmt.Errorf("empty")
	}
	*s = (*s)[0 : l-1]
	return nil
}

func (s *Stack) Push(item Item) error {
	l := len(*s)
	if l == 0 {
		if !item.IsOpen() {
			return NewErr('o', item)
		}
		*s = append(*s, item)
		return nil
	}

	last := (*s)[l-1]
	expect := pairMap[last]

	if !item.IsOpen() {
		if item != expect {
			return NewErr(expect, item)
		}
		s.Pop()
		return nil
	}

	*s = append(*s, item)
	return nil
}

func (s *Stack) PushReverse(item Item) *ItemErr {
	l := len(*s)
	if l == 0 {
		if item.IsOpen() {
			return NewErr('c', item)
		}
		*s = append(*s, item)
		return nil
	}

	last := (*s)[l-1]
	expect := pairMap[last]

	if item.IsOpen() {
		if item != expect {
			return NewErr(expect, item)
		}
		s.Pop()
		return nil
	}

	*s = append(*s, item)
	return nil
}

func (s *Stack) Reverse() *Stack {
	n := make(Stack, len(*s))
	c := 0
	for i := len(*s) - 1; i >= 0; i-- {
		n[c] = (*s)[i]
		c++
	}
	return &n
}
