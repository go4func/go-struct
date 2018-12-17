package main

import (
	"log"
)

func main() {
	values := []int{2, 3, 1, 3, 46, 7}
	Sort(values)
	log.Println(values)
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var t *tree
	for _, value := range values {
		t = add(t, value)
	}
	appendValues(values[:0], t)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func add(t *tree, value int) *tree {
	if t == nil {
		return &tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
