package main

import (
	"fmt"
)

type Node struct {
	next *Node
	value interface{}
}

type LinkedList struct {
	head *Node
}

func (LL *LinkedList) Push(data interface{}){
	new_node := Node{
		next: LL.head,
		value: data,
	}
	LL.head = &new_node
}

func (LL *LinkedList) Insert(prev_node *Node, data interface{}){
	if prev_node != nil {
		new_node :=  Node{
			next:  prev_node.next,
			value: data,
		}
		prev_node.next = &new_node
	}
}

func (LL *LinkedList) Append(data interface{}){
	new_node := Node{
		next:  nil,
		value: data,
	}
	if LL.head != nil {
		tail := LL.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = &new_node
	}else {
		LL.head = &new_node
	}
}

func (LL *LinkedList) Len() int{
	len := 0
	tail := LL.head
	for tail.next != nil {
		len += 1
		tail = tail.next
	}
	return len
}

func (LL *LinkedList) Delete(key interface{}){
	if key == LL.head.value {
		LL.head = LL.head.next
		return
	}
	temp := LL.head
	for temp != nil {
		if temp.next != nil {
			if temp.next.value == key{
				prev_node := temp
				temp = temp.next
				prev_node.next = temp.next
				break
			}
			temp = temp.next
		}else {
			break
		}
	}
}

func (LL *LinkedList) Reverse(){
	if LL.head != nil {
		var prev_node *Node
		curr_node := LL.head
		for curr_node != nil {
			next_node := curr_node.next
			curr_node.next = prev_node
			prev_node = curr_node
			curr_node = next_node
		}
		LL.head = prev_node
	}
}

func (LL *LinkedList) ReverseRecursive() {
	curr_node := LL.head
	if curr_node == nil || curr_node.next == nil {
		return
	}
	LL.head = curr_node.next
	LL.ReverseRecursive()
	curr_node.next.next = curr_node
	curr_node.next = nil
}

func (LL *LinkedList) Print(){
	temp := LL.head
	for temp != nil {
		fmt.Print(temp.value)
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
}

func main() {
	ll := LinkedList{}
	ll.Append(6)
	ll.Len()
	ll.Push(7)
	ll.Push(1)
	ll.Append(4)
	ll.Append(9)
	ll.Insert(ll.head.next, 8)
	ll.Print()
	fmt.Println()
	ll.Delete(1)
	ll.Print()
	ll.ReverseRecursive()
	fmt.Println()
	ll.Print()
	ll.Reverse()
	fmt.Println()
	ll.Print()
	fmt.Println(ll.head)
}