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

func ReadCharsToInts(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	ns := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		tmp := make([]int, 0)
		for _, v := range []rune(line) {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			tmp = append(tmp, n)
		}
		ns = append(ns, tmp)
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

func Avg(ns []int) int {
	var a int
	for _, n := range ns {
		a += n
	}
	return int(a / len(ns))
}

func Sum(ns []int) int {
	var a int
	for _, n := range ns {
		a += n
	}
	return a
}

func Uniq(ns []int) []int {
	m := make(map[int]bool, 0)
	u := make([]int, 0)
	for _, n := range ns {
		m[n] = true
	}
	for k, _ := range m {
		u = append(u, k)
	}
	sort.Ints(u)
	return u
}

// Median returns a median from sorted []int
func Median(ns []int) int {
	l := len(ns)

	if l%2 == 0 {
		return (ns[l/2-1] + ns[l/2]) / 2
	}

	return ns[len(ns)/2]
}

func FindIndex(n int, ns []int) int {
	for i, a := range ns {
		if a == n && ns[i+1] != n {
			return i
		}
	}
	return -1
}
