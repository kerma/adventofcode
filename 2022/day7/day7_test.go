package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/exp/slices"

	. "kerma/adventofcode/2022/aoc"
)

var example = []byte(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`)

type File struct {
	name string
	size int
	dir  *Dir
}

type Dir struct {
	key    string
	size   int
	parent *Dir
}

type Directories struct {
	keys  map[string]struct{}
	items []*Dir
}

func (ds *Directories) Append(d *Dir) {
	if ds.keys == nil {
		ds.keys = make(map[string]struct{})
	}
	if _, ok := ds.keys[d.key]; !ok {
		ds.keys[d.key] = struct{}{}
		ds.items = append(ds.items, d)
	}
}

func (d *Dir) Add(i int) {
	d.size = d.size + i
}

func command(curdir *Dir, s string) (*Dir, *File) {
	if strings.HasPrefix(s, "dir ") || strings.HasPrefix(s, "$ ls") { // ignore
		return curdir, nil
	}

	if strings.HasPrefix(s, "$ cd ..") {
		return curdir.parent, nil
	}

	if strings.HasPrefix(s, "$ cd") {
		var d, key string
		if _, err := fmt.Sscanf(s, "$ cd %s", &d); err != nil {
			panic("invalid cd: " + s)
		}
		if curdir != nil {
			key = curdir.key + "/" + d
		}
		return &Dir{key: key, parent: curdir}, nil
	}

	file := &File{dir: curdir}
	if _, err := fmt.Sscanf(s, "%d %s", &file.size, &file.name); err != nil {
		panic("invalid command: " + s)
	}
	return curdir, file
}

func smallest(dirs []*Dir) int {
	i, root := Last(dirs)
	dirs = dirs[:i]

	total := 70000000
	required := 30000000
	unused := total - root.size

	dirs = Filter(dirs, func(d *Dir) bool {
		return unused+d.size >= required
	})

	slices.SortFunc(dirs, func(a, b *Dir) bool {
		return a.size < b.size
	})

	return dirs[0].size
}

func totals(dirs []*Dir) int {
	dirs = Filter(dirs, func(a *Dir) bool {
		return a.size < 100000
	})
	return SumBy(dirs, func(d *Dir) int {
		return d.size
	})
}

func run(r io.Reader, find func([]*Dir) int) int {
	s := bufio.NewScanner(r)

	var (
		curdir = new(Dir)
		ds     = new(Directories)
		file   = new(File)
		fs     = make([]File, 0)
	)
	for s.Scan() {
		curdir, file = command(curdir, s.Text())
		ds.Append(curdir)
		if file != nil {
			fs = append(fs, *file)
		}
	}

	for _, f := range fs {
		f.dir.Add(f.size)
	}

	dirs := ds.items
	slices.SortFunc(dirs, func(a *Dir, b *Dir) bool {
		return a.key > b.key
	})

	for _, d := range dirs {
		if d.parent != nil {
			d.parent.Add(d.size)
		}
	}

	return find(dirs)
}

func TestExample(t *testing.T) {
	expect := 95437
	r := bytes.NewReader(example)
	if s := run(r, totals); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestTotals(t *testing.T) {
	expect := 1648397

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, totals); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestExampleSmallest(t *testing.T) {
	expect := 24933642
	r := bytes.NewReader(example)
	if s := run(r, smallest); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}

func TestSmallest(t *testing.T) {
	expect := 1815525

	f, _ := os.Open("input")
	defer f.Close()
	if s := run(f, smallest); s != expect {
		t.Fatalf("%d != %d\n", s, expect)
	}
}
