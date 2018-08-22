package main

import (
	"fmt"
	"log"
)

type bstree struct {
	root node
}

type node struct {
	id    int
	left  *node
	right *node
}

func build() bstree {
	arr := []int{4, 3, 2, 5, 8, 9, 22, 1, 7}
	bst := bstree{}
	for i := 0; i < len(arr); i++ {
		apend(arr[i], &bst)
	}
	return bst
}

func apend(v int, bst *bstree) {
	if bst.root.id == 0 {
		log.Println("root empty")
		buildRoot(v, bst)
	} else {
		log.Println("append child")
		appendChild(v, &bst.root)
	}
}

func appendChild(v int, nd *node) {
	if v < nd.id {
		log.Println("append child left", v)
		if nd.left.id == 0 {
			nd.left = &node{v, &node{}, &node{}}
		} else {
			appendChild(v, nd.left)
		}
	} else {
		log.Println("append child right", v)
		if nd.right.id == 0 {
			nd.right = &node{v, &node{}, &node{}}
		} else {
			appendChild(v, nd.right)
		}
	}
}

func buildRoot(v int, bst *bstree) {
	rt := node{v, &node{}, &node{}}
	bst.root = rt
	log.Println("build root")
}

func show(bst bstree) {
	log.Println("show")
	showChild(&bst.root)
}

func showChild(nd *node) {
	if nd.id != 0 {
		log.Println("show child", nd.id)
		if nd.left.id != 0 {
			// log.Println("show child left", nd.left.id)
			showChild(nd.left)
		}
		if nd.right.id != 0 {
			// log.Println("show child right", nd.right.id)
			showChild(nd.right)
		}
	}
}

func find(v int, bst bstree) {
	res := 0
	res = findnd(v, &bst.root, res)
	fmt.Println(res)
}

func findnd(v int, nd *node, res int) int {
	if nd.id == 0 {
		res = 0
		return res
	} else if nd.id > v {
		return findnd(v, nd.left, res)
	} else if nd.id < v {
		return findnd(v, nd.right, res)
	} else {
		res = nd.id
		return res
	}
}

func main() {
	fmt.Println("main")
	bst := build()
	show(bst)
	find(22, bst)
}
