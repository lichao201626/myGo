package main

import (
	"fmt"
	"os"
	"strconv"
)

type line struct {
	Values []int64
	Lineno int64
}

type yhsj struct {
	Lines []line
}

func pL(l line) {
	for _, v := range l.Values {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}

func gnL(l line) line {
	lineno := int64(l.Lineno + 1)
	values := make([]int64, lineno)
	res := line{
		Values: values,
		Lineno: lineno,
	}

	values[0] = 1
	if lineno == 2 {
		values[1] = 1
		return res
	}
	for i := 0; i < len(l.Values)-1; i++ {
		values[i+1] = l.Values[i] + l.Values[i+1]
	}
	values[len(l.Values)] = 1
	return res
}

func gyhsj(x int) yhsj {
	lines := make([]line, x)
	res := yhsj{lines}

	if x < 1 || x > 5 {
		return res
	}

	v := []int64{1}
	l := line{v, 1}
	lines[0] = l
	res.Lines = lines
	if x == 1 {
		return res
	}

	pline := l
	for i := 2; i <= x; i++ {
		nline := gnL(pline)
		res.Lines[i-1] = nline
		pline = nline
	}
	return res
}

func pyhsj(y yhsj) {
	for _, v := range y.Lines {
		pL(v)
	}
}

func main() {
	p := os.Args[1]
	fmt.Println(p)
	x, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	yh := gyhsj(x)
	pyhsj(yh)
}
