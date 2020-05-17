package main

import (
	"fmt"
	"testing"
)

func TestArrayDeque(t *testing.T) {
	arrayDeque := ArrayDeque{}
	_, _ = arrayDeque.Add(1)
	_, _ = arrayDeque.Add(2)
	_, _ = arrayDeque.Add(3)
	_, _ = arrayDeque.Add(4)
	_, _ = arrayDeque.Add(5)
	_, _ = arrayDeque.Add(6)
	_, _ = arrayDeque.Add(7)
	_, _ = arrayDeque.Add(8)
	_, _ = arrayDeque.Add(9)
	_, _ = arrayDeque.Add(10)
	_, _ = arrayDeque.Add(11)

	fmt.Print(arrayDeque.GetLast())
	fmt.Print(arrayDeque.RemoveLast())
	fmt.Print(arrayDeque.GetLast())
}
