package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

var (
	example = []byte(`A Y
B X
C Z
`)
)

type Shape int

type Result int

const (
	Rock Shape = iota + 1
	Paper
	Scissors

	Draw Result = iota + 1
	Lost
	Won
)

var (
	shapes = map[string]Shape{
		"A": Rock,
		"X": Rock,
		"B": Paper,
		"Y": Paper,
		"C": Scissors,
		"Z": Scissors,
	}
	results = map[string]Result{
		"X": Lost,
		"Y": Draw,
		"Z": Won,
	}
	needs = map[Result]map[Shape]Shape{
		Draw: {
			Rock:     Rock,
			Paper:    Paper,
			Scissors: Scissors,
		},
		Lost: {
			Rock:     Scissors,
			Paper:    Rock,
			Scissors: Paper,
		},
		Won: {
			Rock:     Paper,
			Paper:    Scissors,
			Scissors: Rock,
		},
	}
	shapeScore = map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}
	winScore = map[Shape]map[Shape]int{
		Rock: {
			Rock:     3,
			Paper:    0,
			Scissors: 6,
		},
		Paper: {
			Rock:     6,
			Paper:    3,
			Scissors: 0,
		},
		Scissors: {
			Rock:     0,
			Paper:    6,
			Scissors: 3,
		},
	}
)

func shape(x string) Shape {
	s, ok := shapes[x]
	if !ok {
		panic(x)
	}
	return s
}

func result(x string) Result {
	s, ok := results[x]
	if !ok {
		panic(x)
	}
	return s
}

func sum(ns []int) int {
	var a int
	for _, n := range ns {
		a += n
	}
	return a
}

func calc(x, y string) int {
	a := shape(x)
	b := shape(y)
	ss, ok := shapeScore[b]
	if !ok {
		panic("wtf")
	}
	w, ok := winScore[b][a]
	if !ok {
		panic("wtf")
	}
	return ss + w
}

func strat(x, y string) int {
	a := shape(x)
	b := result(y)
	s, ok := needs[b][a]
	if !ok {
		panic("wtf")
	}
	w, ok := winScore[s][a]
	if !ok {
		panic("wtf")
	}
	ss, ok := shapeScore[s]
	if !ok {
		panic("wtf")
	}
	return ss + w
}

func score(r io.Reader, calc func(string, string) int) int {
	rounds := make([]int, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		p := strings.Split(s.Text(), " ")
		rounds = append(rounds, calc(p[0], p[1]))
	}
	return sum(rounds)
}

func TestExample(t *testing.T) {
	expect := 15

	r := bytes.NewReader(example)
	if s := score(r, calc); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestScore(t *testing.T) {
	expect := 14375

	f, _ := os.Open("input")
	if s := score(f, calc); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestExampleStrategy(t *testing.T) {
	expect := 12

	r := bytes.NewReader(example)
	if s := score(r, strat); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestStrategy(t *testing.T) {
	expect := 10274

	f, _ := os.Open("input")
	if s := score(f, strat); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
