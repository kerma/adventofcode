package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	example = []byte(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`)
)

type board [][]int

func run(r io.Reader) int {

	var nums []int
	boards := make([]board, 0)
	b := make(board, 0)

	// read and parse input
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			for _, s := range strings.Split(line, ",") {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				nums = append(nums, n)
			}
			continue
		}

		if len(line) == 0 {
			if len(b) > 0 {
				boards = append(boards, b)
			}
			b = make(board, 0)
			continue
		}

		var ns []int
		line = strings.TrimLeft(line, " ")
		line = strings.ReplaceAll(line, "  ", " ")
		for _, s := range strings.Split(line, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			ns = append(ns, n)
		}
		b = append(b, ns)
	}
	boards = append(boards, b)

	// transform columns to rows
	for ii, b := range boards {
		if len(b) == 0 {
			continue
		}
		for i := 0; i < len(b[0]); i++ {
			var n []int
			for _, row := range b {
				n = append(n, row[i])
			}
			boards[ii] = append(boards[ii], n)
		}
		fmt.Printf("B%d %v\n", ii, boards[ii])
	}

	// play the bingo
	winnerIdx := -1
	lastN := 0
	remove := func(s []int, i int) []int {
		return append(s[:i], s[i+1:]...)
	}
	for _, n := range nums {
		for bi, b := range boards {
			for ri, row := range b {
				for ni, bn := range row {
					if n == bn {
						nrow := remove(row, ni)
						boards[bi][ri] = nrow
						if len(nrow) == 0 {
							winnerIdx = bi
							lastN = n
							goto BINGO
						}
					}
				}
			}
		}
	}

BINGO:

	flatten := func(ns [][]int) []int {
		n := make([]int, 0)
		for _, row := range ns {
			for _, v := range row {
				n = append(n, v)
			}
		}
		return n
	}
	numbers := flatten(boards[winnerIdx])

	filter := func(ns []int) []int {
		contains := func(ns []int, n int) bool {
			for _, v := range ns {
				if v == n {
					return true
				}
			}
			return false
		}

		f := make([]int, 0)
		for _, n := range ns {
			if !contains(f, n) {
				f = append(f, n)
			}
		}
		return f
	}
	numbers = filter(numbers)

	sum := func(ns []int) (r int) {
		for _, n := range ns {
			r += n
		}
		return
	}
	total := sum(numbers) - lastN

	return total * lastN
}

func TestExample(t *testing.T) {
	expect := 4512
	r := bytes.NewReader(example)
	if got := run(r); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}

func TestPartOne(t *testing.T) {
	expect := 39984

	file, err := os.Open("input")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if got := run(file); expect != got {
		t.Fatalf("%d != %d\n", expect, got)
	}
}
