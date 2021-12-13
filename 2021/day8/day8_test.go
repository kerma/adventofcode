package day7

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/kerma/adventofcode/2021/util"
)

var (
	example = []byte(`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`)
)

func read(r io.Reader) ([][]string, [][]string) {

	var patterns, outputs [][]string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " | ")
		if len(parts) != 2 {
			continue
		}

		var t1 []string
		for _, x := range strings.Split(parts[0], " ") {
			t1 = append(t1, x)
		}
		patterns = append(patterns, t1)

		var t2 []string
		for _, x := range strings.Split(parts[1], " ") {
			t2 = append(t2, x)
		}
		outputs = append(outputs, t2)
	}
	return patterns, outputs
}

func filter(d [][]string) [][]string {
	var f [][]string
	for _, line := range d {
		var tmp []string
		for _, p := range line {
			if len(p) != 5 && len(p) != 6 {
				tmp = append(tmp, p)
			}
		}
		f = append(f, tmp)
	}
	return f
}

func run(r io.Reader) int {
	_, patterns := read(r)
	patterns = filter(patterns)

	var c int
	for _, line := range patterns {
		c += len(line)
	}
	return c
}

type Set map[rune]bool

func (s Set) Intersect(s2 Set) Set {
	inter := map[rune]bool{}
	for a, _ := range s {
		if s2[a] {
			inter[a] = true
		}
	}
	return inter
}

func NewSet(str string) Set {
	set := make(Set)
	for _, r := range str {
		set[r] = true
	}
	return set
}

func run2(r io.Reader) int {

	lines, outputs := read(r)

	values := make([]int, len(lines))
	for i, signals := range lines {
		knownSets := make(map[int]Set)
		for _, sig := range signals {
			if len(sig) == 2 || len(sig) == 4 {
				knownSets[len(sig)] = NewSet(sig)
			}
		}

		decoded := ""
		for _, output := range outputs[i] {
			switch len(output) {
			case 2:
				decoded += "1"
			case 3:
				decoded += "7"
			case 4:
				decoded += "4"
			case 7:
				decoded += "8"
			case 5:
				s := NewSet(output)
				if len(s.Intersect(knownSets[2])) == 2 {
					decoded += "3"
				} else if len(s.Intersect(knownSets[4])) == 2 {
					decoded += "2"
				} else {
					decoded += "5"
				}
			default: // len == 6
				s := NewSet(output)
				if len(s.Intersect(knownSets[2])) == 1 {
					decoded += "6"
				} else if len(s.Intersect(knownSets[4])) == 4 {
					decoded += "9"
				} else {
					decoded += "0"
				}
			}
		}
		v, _ := strconv.Atoi(decoded)
		values[i] = v
	}

	return util.Sum(values)
}

func TestExample(t *testing.T) {
	expect := 26
	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 554

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
	expect := 61229

	r := bytes.NewReader(example)
	if got := run2(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 990964

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run2(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
