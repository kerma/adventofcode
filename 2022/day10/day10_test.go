package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	. "kerma/adventofcode/2022/aoc"
)

var example = []byte(`addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`)

type CPU struct {
	x           int
	cycle       int
	breakpoints map[int]struct{}
	store       []int
}

func NewCPU(bp []int) *CPU {
	m := make(map[int]struct{})
	for _, n := range bp {
		m[n] = struct{}{}
	}
	return &CPU{
		x:           1,
		cycle:       0,
		breakpoints: m,
		store:       make([]int, 0, len(bp)),
	}
}

func (cpu *CPU) exec() {
	cpu.cycle++
	if _, ok := cpu.breakpoints[cpu.cycle]; ok {
		cpu.store = append(cpu.store, cpu.cycle*cpu.x)
	}
}

func (cpu *CPU) Run(inst string, x int) {
	cpu.exec()
	if inst == "addx" {
		cpu.exec()
		cpu.x = cpu.x + x
	}
}

func run(r io.Reader) int {
	s := bufio.NewScanner(r)

	cpu := NewCPU([]int{20, 60, 100, 140, 180, 220})
	for s.Scan() {
		line := s.Text()
		if line == "noop" {
			cpu.Run(line, 0)
			continue
		}
		parts := strings.Split(line, " ")
		n, _ := strconv.Atoi(parts[1])
		cpu.Run(parts[0], n)
	}

	return Sum(cpu.store)
}

func TestExample(t *testing.T) {
	expect := 13140
	r := bytes.NewReader(example)
	if s := run(r); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestVisible(t *testing.T) {
	expect := 14860

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
