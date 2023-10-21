package linkedlist

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
}

type List[T any] struct {
	len  int
	head *node[T]
}

func NewSingly[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) InsertHead(data T) {
	newNode := &node[T]{data: data, next: nil}
	newNode.next = l.head
	l.head = newNode
	l.len++
}

func (l *List[T]) InsertTail(data T) {
	if l.len == 0 {
		l.InsertHead(data)
		return
	}

	l.add(data)
}

func (l *List[T]) Insert(data T, index int) error {
	if index < 0 {
		return IndexOutOfRangeError
	}

	if index > l.len {
		return IndexOutOfRangeError
	}

	if index == 0 {
		l.InsertHead(data)
		return nil
	}

	if index == l.len {
		l.add(data)
		return nil
	}

	i := 1
	curNode := l.head
	prevNode := curNode

	for curNode != nil {
		if i == index {
			newNode := &node[T]{data: data, next: nil}
			prevNode.next = newNode
			newNode.next = curNode
			l.len++
			return nil
		}

		i++
		prevNode = curNode
		curNode = curNode.next
	}

	return UnknownError
}

func (l *List[T]) RemoveHead() error {
	if l.len == 0 {
		return ListIsEmptyError
	}

	l.head = l.head.next
	l.len--

	return nil
}

func (l *List[T]) RemoveTail() error {
	if l.len == 0 {
		return ListIsEmptyError
	}

	if l.len == 1 {
		err := l.RemoveHead()
		if err != nil {
			return err
		}
	}

	curNode := l.head
	for curNode.next.next != nil {
		curNode = curNode.next
	}

	curNode.next = nil
	l.len--

	return nil
}

func (l *List[T]) Remove(index int) error {
	if index < 0 {
		return IndexOutOfRangeError
	}

	if index >= l.len {
		return IndexOutOfRangeError
	}

	if index == 0 {
		return l.RemoveHead()
	}

	i := 1
	curNode := l.head

	for curNode != nil {
		if i == index {
			curNode.next = curNode.next.next
			l.len--
			return nil
		}

		i++
		curNode = curNode.next
	}

	return UnknownError
}

func (l *List[T]) Head() T {
	return l.head.data
}

func (l *List[T]) Tail() T {
	curNode := l.head
	for curNode.next != nil {
		curNode = curNode.next
	}

	return curNode.data
}

func (l *List[T]) Transverse() {
	for curNode := l.head; curNode != nil; curNode = curNode.next {
		fmt.Print(curNode.data, " ")
	}
	fmt.Println()
}

func (l *List[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *List[T]) Length() int {
	return l.len
}

func (l *List[T]) add(data T) {
	curNode := l.head
	for curNode.next != nil {
		curNode = curNode.next
	}

	curNode.next = &node[T]{data: data, next: nil}
	l.len++
}
