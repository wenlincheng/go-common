package collection

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.AddFirst(4)

	err, item := linkedList.GetLast()
	fmt.Println(item)

	err, item1 := linkedList.RemoveLast()
	fmt.Println(item1)

	err, item2 := linkedList.GetLast()
	fmt.Println(item2)

	contain := linkedList.Contains(1)
	fmt.Println(contain)
	contain2 := linkedList.Contains(3)
	fmt.Println(contain2)
	err, item3 := linkedList.Get(4)
	fmt.Printf("Size %d \n", linkedList.Size())

	fmt.Printf("%s 节点值：%v", err, item3)
}
