package main

import "github.com/kerma/adventofcode/v21/util"

type submarine interface {
	forward(int) submarine
	down(int) submarine
	up(int) submarine
	abs() int
}

type pos struct {
	x int
	y int
}

func (p pos) forward(n int) submarine {
	return pos{
		p.x + n,
		p.y,
	}
}

func (p pos) down(n int) submarine {
	return pos{
		p.x,
		p.y - n,
	}
}

func (p pos) up(n int) submarine {
	return pos{
		p.x,
		p.y + n,
	}
}

func (p pos) abs() int {
	return util.Abs(p.x * p.y)
}

type posV2 struct {
	x int
	y int
	z int
}

func (p posV2) forward(n int) submarine {
	return posV2{
		p.x + n,
		p.y + (p.z * n),
		p.z,
	}
}

func (p posV2) down(n int) submarine {
	return posV2{
		p.x,
		p.y,
		p.z + n,
	}
}

func (p posV2) up(n int) submarine {
	return posV2{
		p.x,
		p.y,
		p.z - n,
	}
}

func (p posV2) abs() int {
	return util.Abs(p.x * p.y)
}
