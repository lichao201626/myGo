package main

import (
	"fmt"
)

type btree struct {
	root node
}

type node struct {
	id    int
	left  *node
	right *node
}

func show(bt btree) {
	fmt.Printf("show")
	dep := 0
	for true {
		f := getf(bt, dep)
		fmt.Println()
		for i := 0; i < len(f); i++ {
			fmt.Printf("\t%v", f[i].id)
		}
		n := isfull(f)
		if n.id != 0 {
			break
		}
		dep++
	}
	fmt.Println()
	f := getf(bt, dep+1)
	for i := 0; i < len(f); i++ {
		if f[i].id != 0 {
			fmt.Printf("\t%v", f[i].id)
		}
	}
}

func build() btree {
	// {5, 4, 6, 8, 0, 9, 3, 2, 1}
	arr := []int{5, 4, 6, 8, 11, 9, 3, 2, 1}
	bt := btree{}
	buildRoot(arr[0], &bt)
	for i := 1; i < len(arr); i++ {
		p := pos(i + 1)
		tmp := &bt.root
		for j := len(p) - 1; j > 0; j-- {
			if p[j] == "right" {
				tmp = tmp.right
			}
			if p[j] == "left" {
				tmp = tmp.left
			}
		}
		// append(arr[i], &bt)
		create(arr[i], tmp)
	}
	return bt
}

func buildRoot(v int, bt *btree) {
	right := node{}
	left := node{}
	rt := node{v, &left, &right}
	bt.root = rt
}

func pos(pos int) []string {
	res := make([]string, pos)
	i := 0
	for pos > 1 {
		s := pos % 2
		if s == 0 {
			res[i] = "left"
		} else {
			res[i] = "right"
		}
		pos = pos / 2
		i++
	}
	return res
}

// depth first append
func append(v int, bt *btree) {
	dep := 0
	for true {
		f := getf(*bt, dep)
		n := isfull(f)

		if n.id != 0 {
			create(v, &bt.root)
			break
		}
		dep++
	}
}

func getf(bt btree, dep int) []node {
	// res := make([]node, 2^dep)
	temp := getc(bt, []node{}, 0)
	for i := 1; i <= dep; i++ {
		temp = getc(bt, temp, i)
	}
	return temp
}

func getc(bt btree, pref []node, dep int) []node {
	x := 1
	for i := 0; i < dep; i++ {
		x *= 2
	}
	res := make([]node, x)
	if len(pref) == 0 {
		res[0] = bt.root
	}
	for i := 0; i < len(pref); i++ {
		res[2*i] = *pref[i].left
		res[2*i+1] = *pref[i].right
	}

	return res
}

func isfull(par []node) *node {
	res := node{}
	for i := 0; i < len(par); i++ {
		if par[i].left.id == 0 || par[i].right.id == 0 {
			res = par[i]
			break
		}
	}
	return &res
}

func create(v int, nd *node) {
	if nd.left.id == 0 {
		// fmt.Printf("add left %v", v)
		nd.left = &node{v, &node{}, &node{}}
	} else {
		// fmt.Printf("add right %v", v)
		nd.right = &node{v, &node{}, &node{}}
	}
}

func main() {
	bt := build()
	show(bt)
}
