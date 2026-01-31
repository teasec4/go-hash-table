package main

import (
	"cmp"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

// Node
type Node[T any] struct{
	val T
	left, right *Node[T]
}

type Tree[T any] struct{
	f OrderableFunc[T]
	root *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T]{
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(v T){
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool{
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T]{
	if n == nil{
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val);{
		case r <= -1:
		n.left = n.left.Add(f, v)
		case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool{
	if n == nil{
		return false
	}
	switch r:= f(v, n.val);{
		case r<= -1:
		return n.left.Contains(f, v)
		case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

func main() {
	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(15)
	t1.Add(8)
	
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(20))
	
	v1 := 1
	v2 := 2.5
	v1 = doubleValue(v1)
	v2 = doubleValue(v2)
	fmt.Println(v1)
	fmt.Println(v2)
	
}


type CustomNumber interface{
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

func doubleValue[T CustomNumber](value T) T{
	return value * 2
}

type Printable interface{
	fmt.Stringer
	~int | ~float64
}

func PrintValue[T Printable](v T) {
	fmt.Println(v.String())
}