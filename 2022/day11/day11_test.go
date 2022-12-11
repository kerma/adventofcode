package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

var example = []byte(`Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1

`)

type Monkey struct {
	items      []int64
	op         func(int64) int64
	test       int64
	success    int
	fail       int
	inspected  int
	divProduct int64 // product of every test value, 69577 for the example
}

func (m *Monkey) String() string {
	return fmt.Sprintf("%v", *m)
}

func runRoundPart1(m *Monkey, ms []*Monkey) {
	for _, item := range m.items {
		w := int64(m.op(item) / 3)
		if w%m.test == 0 {
			target := ms[m.success]
			target.items = append(target.items, w)
		} else {
			target := ms[m.fail]
			target.items = append(target.items, w)
		}
		m.inspected++
	}
	m.items = []int64{}
}

func runRoundPart2(m *Monkey, ms []*Monkey) {
	for _, item := range m.items {
		w := m.op(item)
		if w%m.test == 0 {
			target := ms[m.success]
			target.items = append(target.items, w%m.divProduct)
		} else {
			target := ms[m.fail]
			target.items = append(target.items, w%m.divProduct)
		}
		m.inspected++
	}
	m.items = []int64{}
}

type Op func(int64) int64

func Add(x int64) Op {
	return func(i int64) int64 {
		return x + i
	}
}

func Mul(x int64) Op {
	return func(i int64) int64 {
		return x * i
	}
}

func Square(x int64) int64 {
	return x * x
}

func run(r io.Reader, rounds int, runRound func(*Monkey, []*Monkey)) int {
	monkeys := []*Monkey{}

	m := &Monkey{}
	divProduct := int64(1)
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "  Starting items: ") {
			p := strings.Split(line, ": ")
			var items []int64
			for _, i := range strings.Split(p[1], ", ") {
				n, _ := strconv.ParseInt(i, 10, 0)
				items = append(items, n)
			}
			m.items = items
		}
		if strings.HasPrefix(line, "  Operation: ") {
			var op string
			var n int64
			if _, err := fmt.Sscanf(line, "  Operation: new = old %s %d", &op, &n); err == nil {
				if op == "*" {
					m.op = Mul(n)
				} else {
					m.op = Add(n)
				}
			} else {
				m.op = Square
			}
		}
		if strings.HasPrefix(line, "  Test: ") {
			var n int64
			fmt.Sscanf(line, "  Test: divisible by %d", &n)
			m.test = n
			divProduct = divProduct * n
		}
		if strings.HasPrefix(line, "    If true") {
			var n int
			fmt.Sscanf(line, "    If true: throw to monkey %d", &n)
			m.success = n
		}
		if strings.HasPrefix(line, "    If false") {
			var n int
			fmt.Sscanf(line, "    If false: throw to monkey %d", &n)
			m.fail = n
		}
		if line == "" {
			monkeys = append(monkeys, m)
			m = &Monkey{}
		}
	}

	for _, m := range monkeys {
		m.divProduct = divProduct
	}

	for r := 1; r <= rounds; r++ {
		for _, m := range monkeys {
			runRound(m, monkeys)
		}
	}

	slices.SortFunc(monkeys, func(a, b *Monkey) bool {
		return a.inspected > b.inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}

func TestExample(t *testing.T) {
	expect := 10605
	r := bytes.NewReader(example)
	if s := run(r, 20, runRoundPart1); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestWorry(t *testing.T) {
	expect := 95472

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, 20, runRoundPart1); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestBigExample(t *testing.T) {
	expect := 2713310158
	r := bytes.NewReader(example)
	if s := run(r, 10000, runRoundPart2); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestBig(t *testing.T) {
	expect := 17926061332

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, 10000, runRoundPart2); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
