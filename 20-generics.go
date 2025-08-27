package main

//Generics are type parameters

//SlicesIndex takes a slice of any comparable type and an element of that type and returns the index of the first occurence
func SlicesIndex[S ~ []E, E comparable] (s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

//List is a linked list with values of any type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

//Define methods on generic types. The type is List[T] not List

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

//AllElements returns elementsin List as a slice. Iterate over all elements of custom types

func (lst *List[T]) AllElements() []T {
	var elements []T 
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}