package main

import ( 
	"fmt"
)


// Node
type Node[T any] struct{
	val T
	next *Node[T]
}

type List[T any] struct{
	Head *Node[T]
	Tail *Node[T]
	Len int 
}

func (l *List[T]) Add(v T){
	newNode := &Node[T]{val:v}
	
	if l.Head == nil{
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.next = newNode
		l.Tail = newNode
	}
	
	l.Len++
}

func (l *List[T]) Get(index int) (T, bool){
	var zero T
	
	if index < 0 || index >=l.Len {
		return zero, false
	}
	
	cur := l.Head
	for i := 0; i < index; i++{
		cur = cur.next
	}
	
	return cur.val, true
}

func (l *List[T]) Print() {
	for n := l.Head; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}

func (l *List[T]) Insert(index int, v T) bool{
	if index < 0 || index > l.Len{
		return false
	}
	
	node := &Node[T]{val: v}
	
	if index == 0{
		node.next = l.Head
		l.Head = node
		if l.Len == 0 {
			l.Tail = node
		}
		l.Len++
		return true
	}
	
	if index == l.Len{
		l.Tail.next = node
		l.Tail = node
		l.Len++
		return true
	}
	
	prev := l.Head
	for i := 0; i < index-1; i++{
		prev = prev.next
	}
	node.next = prev.next
	prev.next = node
	l.Len++
	
	return true
}


func main() {
	var list List[string]

	list.Add("a")
	list.Add("b")
	list.Add("c")

	list.Print()
	
	s, _ := list.Get(1)
	fmt.Println(s)
	list.Insert(1, "m")
	s, _ = list.Get(1)
	fmt.Println(s)
}