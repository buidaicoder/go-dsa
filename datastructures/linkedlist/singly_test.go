package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestList_InsertHead(t *testing.T) {
	list := NewSingly[int]()
	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)

	want := []any{1, 2, 3}
	var got []any
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestList_InsertTail(t *testing.T) {
	list := NewSingly[int]()
	list.InsertTail(4)
	list.InsertTail(5)
	list.InsertTail(6)

	want := []any{4, 5, 6}
	var got []any
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestList_Insert(t *testing.T) {
	list := NewSingly[int]()
	list.InsertHead(1)
	list.InsertTail(2)

	if err := list.Insert(3, 2); err != nil {
		t.Errorf(err.Error())
	}

	want := []any{1, 2, 3}
	var got []any
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestList_RemoveHead(t *testing.T) {
	list := NewSingly[int]()
	_ = list.Insert(1, 0)
	_ = list.Insert(2, 1)
	_ = list.Insert(3, 2)

	if err := list.RemoveHead(); err != nil {
		t.Errorf(err.Error())
	}

	want := []any{2, 3}
	var got []any
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestList_RemoveTail(t *testing.T) {
	list := NewSingly[int]()
	_ = list.Insert(1, 0)
	_ = list.Insert(2, 1)
	_ = list.Insert(3, 2)

	if err := list.RemoveTail(); err != nil {
		fmt.Println(err)
	}

	want := []any{1, 2}
	var got []any
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestList_Remove(t *testing.T) {
	list := NewSingly[int]()
	_ = list.Insert(1, 0)
	_ = list.Insert(2, 1)
	_ = list.Insert(3, 2)

	if err := list.Remove(1); err != nil {
		fmt.Println(err)
	}

	want := []any{1, 3}
	var got []any
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
