package day7

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
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

type data [][]string

func read(r io.Reader) (data, data) {

	var patterns, outputs data

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

func filter(d data) data {
	var f data
	for _, line := range d {
		var tmp []string
		for _, p := range line {
			if len(p) != 1 && len(p) != 5 && len(p) != 6 {
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
	fmt.Printf("%v\n", patterns)

	var c int
	for _, line := range patterns {
		c += len(line)
	}
	return c
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
