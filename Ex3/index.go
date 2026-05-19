package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *List[T]) Prepend(val T) {
	newEl := &element[T]{val: val, next: l.head}
	l.head = newEl
	if l.tail == nil {
		l.tail = newEl
	}
}

func (l *List[T]) Append(val T) {
	newEl := &element[T]{val: val}
	if l.head == nil {
		l.head = newEl
		l.tail = newEl
		return
	}
	l.tail.next = newEl
	l.tail = newEl
}

func (l *List[T]) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *List[T]) Search(value T) bool {
	current := l.head
	for current != nil {
		if current.val == value {
			return true
		}
		current = current.next
	}
	return false
}

func (l *List[T]) Delete(val T, eq func(T, T) bool) {
	if l.head == nil {
		return
	}
	if eq(l.head.val, val) {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		return
	}
	current := l.head
	for current.next != nil && !eq(current.next.val, val) {
		current = current.next
	}
	if current.next != nil {
		if current.next == l.tail {
			l.tail = current
		}
		current.next = current.next.next
	}
}

func main() {
	// ll := LinkedList{}
	// ll.Append(10)
	// ll.Append(20)
	// ll.Append(30)
	// ll.Append(40)
	// ll.Append(50)
	// ll.PrintList() // 10 -> 20 -> 30 -> 40 -> 50 -> nil
	// ll.InsertAt(2, 9999)
	// ll.PrintList()

	// ll.InsertAt(0, 1)
	// ll.PrintList()

	// ll.InsertAt(100, 777)
	// ll.PrintList()

	// Test Generic List
	gl := List[int]{}
	gl.Append(10)
	gl.Append(20)
	gl.Append(30)
	gl.PrintList() // 10 -> 20 -> 30 -> nil

	gl.Prepend(5)
	gl.PrintList() // 5 -> 10 -> 20 -> 30 -> nil

	gl.Delete(20, func(a, b int) bool { return a == b })
	gl.PrintList() // 5 -> 10 -> 30 -> nil

	// Test với string
	gs := List[string]{}
	gs.Append("hello")
	gs.Append("world")
	gs.PrintList() // hello -> world -> nil

}

// EXCERCISE: 1. Viết method InsertAt() để chèn vào bất kỳ vị trí nào mình muốn
func (ll *LinkedList) InsertAt(position int, value int) {
	newNode := &Node{data: value}

	if position == 0 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head
	for i := 0; i < position-1; i++ {
		if current == nil {
			return
		}
		current = current.next
	}

	if current == nil {
		return
	}

	newNode.next = current.next
	current.next = newNode
}

// EXCERCISE: 2. Thay đổi toàn bộ data type thành Generic Types !
type List[T comparable] struct {
	head, tail *element[T]
}

type element[T comparable] struct {
	val  T
	next *element[T]
}