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

type leaf struct {
	id int
}

func show() {
	fmt.Printf("show")
}

func build() btree {
	arr := []int{5, 4, 6, 8, 0, 9, 3, 2, 1}
	bt := btree{}
	buildRoot(arr[0], &bt)
	for i := 1; i < len(arr); i++ {
		append(arr[i], &bt)
	}
	return bt
}

func buildRoot(v int, bt *btree) {
	right := node{}
	left := node{}
	rt := node{v, &left, &right}
	bt.root = rt
}

// depth first append
func append(v int, bt *btree) {
	// get the last parent of a btree
	pre := bt.root.left
	bro := bt.root.right
	if bt.root.left != nil && bt.root.right != nil {
		pre := bt.root.left
	}
	if bt.root.left != nil && bt.root.right == nil {
		// return
	}

}

func getn(bt btree, bro btree) node {
	if bt.root.left != nil && bt.root.right != nil {
		pre := bt.root.left
		if bro.root != nil {
			getn(bro)
		}
	}
	if bt.root.left != nil && bt.root.right == nil {
		// return
		return bt.root
	}
	if bt.root.left == nil {
		// return
		return bt.root
	}
	return
}

func create(v int, nd *node) {
	if nd.left == nil {
		nd.left = &node{v, &node{}, &node{}}
	} else {
		nd.right = &node{v, &node{}, &node{}}
	}
}

func main() {
	build()
	show()
}
