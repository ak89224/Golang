package main

import "fmt"

type deque struct {
	arr []interface{}
	front int
	rear  int
	size  int
}

func NewDeque(size int) *deque {
	var dq *deque
	if size !=0 {
		dq = &deque{
			arr:   make([]interface{}, size),
			front: -1,
			rear:  0,
			size:  size,
		}
	}
	return dq
}

func (dq *deque) IsFull() bool {
	return (dq.front == 0 && dq.rear == dq.size-1) || (dq.front == dq.rear+1)
}

func (dq *deque) IsEmpty() bool {
	return dq.front == -1
}

func (dq *deque) InsertFront(item interface{}) {
	if dq.IsFull(){
		fmt.Println("Queue Overflow !")
		return
	}
	if dq.IsEmpty() {
		dq.front = 0
		dq.rear = 0
	}else {
		if dq.front ==0 {
			dq.front = dq.size -1
		}else {
			dq.front = dq.front - 1
		}
	}
	dq.arr[dq.front] = item
}

func (dq *deque) InsertRear(item interface{}) {
	if dq.IsFull(){
		fmt.Println("Queue Overflow !")
		return
	}
	if dq.IsEmpty() {
		dq.front = 0
		dq.rear = 0
	}else {
		if dq.rear == dq.size-1 {
			dq.rear = 0
		}else {
			dq.rear = dq.rear + 1
		}
	}
	dq.arr[dq.rear] = item
}

func (dq *deque) DeleteFront() {
	if dq.IsEmpty() {
		fmt.Println("Deque UnderFlow !")
		return
	}
	dq.arr[dq.front] = nil
	if dq.front == dq.rear {
		dq.front = -1
		dq.rear = -1
	}else {
		if dq.front == dq.size-1 {
			dq.front = 0
		}else {
			dq.front = dq.front + 1
		}
	}
}

func (dq *deque) DeleteRear() {
	if dq.IsEmpty() {
		fmt.Println("Deque UnderFlow !")
		return
	}
	dq.arr[dq.rear] = nil
	if dq.front == dq.rear {
		dq.front = -1
		dq.rear = -1
	}else {
		if dq.rear == 0 {
			dq.rear = dq.size - 1
		}else {
			dq.rear = dq.rear - 1
		}
	}
}


func (dq *deque) GetFront() interface{}{
	if dq.IsEmpty() {
		fmt.Println("Queue UnderFlow !")
		return nil
	}
	return dq.arr[dq.front]
}

func (dq *deque) GetRear() interface{}{
	if dq.IsEmpty() {
		fmt.Println("Queue UnderFlow !")
		return nil
	}
	return dq.arr[dq.rear]
}

func main() {
	dq := NewDeque(5)

	dq.InsertRear(5)
	dq.InsertFront(4)
	dq.InsertRear(3)
	dq.InsertFront(2)
	dq.InsertRear(1)
	dq.InsertFront(0)
	fmt.Println(dq.GetRear())
	fmt.Println(dq.GetFront())
	fmt.Println(dq.arr)
	dq.DeleteFront()
	dq.DeleteRear()
	fmt.Println(dq.GetRear())
	fmt.Println(dq.GetFront())
	fmt.Println(len(dq.arr))
	fmt.Println(dq.arr)
}


