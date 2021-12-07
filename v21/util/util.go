package util

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInts(r io.Reader) []int {
	scanner := bufio.NewScanner(r)
	ns := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, ',') {
			for _, x := range strings.Split(line, ",") {
				n, err := strconv.Atoi(x)
				if err != nil {
					panic(err)
				}
				ns = append(ns, n)
			}
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ns = append(ns, n)
		}
	}
	return ns
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Median(i []int) int {
	if len(i)%2 != 0 { // is even
		panic("Odd number of ints")
	}

	ns := make([]int, len(i))
	copy(ns, i)
	sort.Ints(ns)

	return ns[len(ns)/2]
}
