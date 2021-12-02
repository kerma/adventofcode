package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/kerma/adventofcode/v21/util"
)

var (
	example = []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
)

func parseMovement(s string) (string, int) {
	parts := strings.Split(s, " ")
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return parts[0], n
}

func calculate(p submarine, lines []string) int {
	for _, line := range lines {
		cmd, n := parseMovement(line)
		switch cmd {
		case "forward":
			p = p.forward(n)
		case "down":
			p = p.down(n)
		case "up":
			p = p.up(n)
		}
	}
	return p.abs()
}

func TestExample(t *testing.T) {
	expect := 150
	if got := calculate(pos{}, example); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPosition(t *testing.T) {
	expect := 1924923
	lines, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}

	if got := calculate(pos{}, lines); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExampleV2(t *testing.T) {
	expect := 900
	if got := calculate(posV2{}, example); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPositionV2(t *testing.T) {
	expect := 1982495697
	lines, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}

	if got := calculate(posV2{}, lines); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
