package day3

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	example = []byte(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)
)

func calculate(r io.Reader, size int) int {

	// rotate lines
	var lines = make([]string, size)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			lines[pos] = lines[pos] + string(char)
		}
	}

	// build rates in binary format
	var gr, er string
	for _, row := range lines {
		if strings.Count(row, "1") > len(row)/2 {
			gr = gr + "1"
			er = er + "0"
		} else {
			gr = gr + "0"
			er = er + "1"
		}
	}

	// parse to ints
	grate, err := strconv.ParseInt(gr, 2, 64)
	if err != nil {
		panic(err)
	}
	erate, err := strconv.ParseInt(er, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(grate * erate)
}

func rating(r [][]rune, ox, co rune) int64 {
	lines := make([][]rune, len(r))
	copy(lines, r)

	var bitCounter int
	for {

		// count bits
		var ones, zeros int
		for _, n := range lines {
			if n[bitCounter] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		// filter lines
		tmp := make([][]rune, 0)
		for _, n := range lines {
			if ones > zeros && n[bitCounter] == ox {
				tmp = append(tmp, n)
			} else if ones < zeros && n[bitCounter] == co {
				tmp = append(tmp, n)
			} else if ones == zeros && n[bitCounter] == ox {
				tmp = append(tmp, n)
			}
		}
		lines = tmp

		if len(lines) <= 1 {
			break
		}
		bitCounter++
	}

	rate, err := strconv.ParseInt(string(lines[0]), 2, 64)
	if err != nil {
		panic(err)
	}
	return rate
}

func lifeSupport(r io.Reader, size int) int {

	var lines = make([][]rune, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}

	ox := rating(lines, '1', '0')
	co := rating(lines, '0', '1')

	return int(ox * co)
}

func TestExample(t *testing.T) {
	expect := 198
	r := bytes.NewReader(example)
	if got := calculate(r, 5); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestCalculate(t *testing.T) {
	expect := 3277364

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := calculate(file, 12); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestExamplePart2(t *testing.T) {
	expect := 230
	r := bytes.NewReader(example)
	if got := lifeSupport(r, 5); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPart2(t *testing.T) {
	expect := 5736383

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := lifeSupport(file, 12); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
