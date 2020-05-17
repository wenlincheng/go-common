package main

import (
	"errors"
)

/*
	ArrayDeque是一个双端队列实现，
	内部使用数组进行元素存储，不允许存储null值，
	可以高效的进行元素查找和尾部插入取出，
	是用作队列、双端队列、栈的绝佳选择，性能比LinkedList还要好。
*/

const (
	// 初始化最小容量
	MinInitialCapacity int = 8
)

type ArrayDeque struct {
	elements []interface{}
	head     int
	tail     int
}

// 返回队列的大小
func (a *ArrayDeque) Size() int {
	return (a.tail - a.head) & (len(a.elements) - 1)
}

// 判断队列的是否为空
func (a *ArrayDeque) IsEmpty() bool {
	return a.tail == a.head
}

func (a *ArrayDeque) Clear() {
	h := a.head
	t := a.tail
	if h != t { // clear all cells
		a.head, a.tail = 0, 0
		mask := cap(a.elements) - 1
		i := h
		for {
			a.elements[i] = nil
			i = (i + 1) & mask
			if i == t {
				break
			}
		}
	}
}

// 将指定的元素插入此双端队列的末尾
func (a *ArrayDeque) Add(e interface{}) (error, bool) {
	err := a.AddLast(e)
	if err != nil {
		return err, false
	}
	return nil, true
}

// 在此双端队列的末尾插入指定的元素
func (a *ArrayDeque) AddLast(e interface{}) error {
	if a.elements == nil {
		a.elements = make([]interface{}, 0, 16)
	}
	if e == nil {
		return errors.New("NullPointer Error")
	}
	// 向数组中末尾添加元素，数组容量不够会自动扩容为原来的两倍
	a.tail = a.tail + 1
	a.elements = append(a.elements, e)
	return nil
}

// 在此双端队列的前面插入指定的元素
func (a *ArrayDeque) AddFirst(e interface{}) error {
	if a.elements == nil {
		a.elements = make([]interface{}, 16, 16)
	}
	if e == nil {
		return errors.New("NullPointer Error")
	}
	// 向数组中头部添加元素，数组容量不够会自动扩容为原来的两倍
	a.head = a.head - 1

	a.elements[(a.head)&(len(a.elements)-1)] = e

	a.elements = append(a.elements, e)
	return nil
}

// 将数组容量扩成之前的两倍
func (a *ArrayDeque) doubleCapacity() error {
	p := a.head
	n := len(a.elements)
	r := n - p // number of elements to the right of p
	newCapacity := n << 1
	if newCapacity < 0 {
		return errors.New("Sorry, deque too big")
	}
	newArray := make([]interface{}, newCapacity, newCapacity)
	// 容量变为原来的两倍，然后把head之后的元素复制到新数组的开头，把剩余的元素复制到新数组之后
	a.arrayCopy(a.elements, p, newArray, 0, r)
	a.arrayCopy(a.elements, 0, newArray, r, p)
	a.elements = newArray
	a.head = 0
	a.tail = n

	return nil
}

// 复制数组
func (a *ArrayDeque) arrayCopy(src []interface{}, srcPos int, dest []interface{}, destPos int, length int) {
	for i := srcPos; i < srcPos+length; i++ {
		dest[destPos] = src[i]
		destPos++
	}
}

// 将指定的元素插入此双端队列的前面
func (a *ArrayDeque) OfferFirst(e interface{}) bool {

	return false
}
func (a *ArrayDeque) OfferLast(e interface{}) bool {
	return false
}

// 删除队首元素
func (a *ArrayDeque) RemoveFirst() (error, interface{}) {
	e := a.PollFirst()
	if e == nil {
		return errors.New("NoSuchElement"), nil
	}
	return nil, e
}

// 删除队尾元素
func (a *ArrayDeque) RemoveLast() (error, interface{}) {
	e := a.PollLast()
	if e == nil {
		return errors.New("NoSuchElement"), nil
	}
	return nil, e
}
func (a *ArrayDeque) PollFirst() interface{} {
	h := a.head
	result := a.elements[h]
	// Element is nil if deque empty
	if result == nil {
		return nil
	}
	a.elements[h] = nil

	a.head = (h + 1) & (cap(a.elements) - 1)
	return result
}
func (a *ArrayDeque) PollLast() interface{} {
	t := (a.tail - 1) & (cap(a.elements) - 1)
	result := a.elements[t]
	if result == nil {
		return nil
	}
	a.elements[t] = nil
	a.tail = t
	return result
}

// 返回列首元素
func (a *ArrayDeque) GetFirst() (error, interface{}) {
	result := a.elements[a.head]
	if result == nil {
		return errors.New("NoSuchElement"), nil
	}
	return nil, result
}

// 返回队尾元素
func (a *ArrayDeque) GetLast() (error, interface{}) {
	result := a.elements[(a.tail-1)&(cap(a.elements)-1)]
	if result == nil {
		return errors.New("NoSuchElement"), nil
	}
	return nil, result
}
func (a *ArrayDeque) PeekFirst() interface{} {
	return nil
}
func (a *ArrayDeque) PeekLast() interface{} {
	return nil
}

func (a *ArrayDeque) Push(e interface{}) {

}
func (a *ArrayDeque) Element() interface{} {
	return nil
}
func (a *ArrayDeque) Peek() interface{} {
	return nil
}
func (a *ArrayDeque) Offer(e interface{}) bool {
	return false
}
func (a *ArrayDeque) Poll() interface{} {
	return nil
}
func (a *ArrayDeque) Pop() interface{} {
	return nil
}
func (a *ArrayDeque) Remove() interface{} {
	return nil
}
