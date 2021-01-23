package main

import (
	"fmt"
)

type Node struct {
	next *Node
	value interface{}
}

type LinkedList struct {
	tail *Node
}

func (LL *LinkedList) Push(data interface{}){
	new_node := Node{
		next:  nil,
		value: data,
	}
	if LL.tail != nil {
		new_node.next = LL.tail.next
		LL.tail.next = &new_node
	}
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
	if LL.tail != nil {
		new_node.next = LL.tail.next
		LL.tail.next = &new_node
		LL.tail = &new_node

	}else {
		LL.tail = &new_node
		new_node.next = LL.tail
	}
}

func (LL *LinkedList) Len() int{
	var len int
	if LL.tail != nil{
		len += 1
		head := LL.tail.next
		for head != LL.tail{
			len += 1
			head = head.next
		}
	}
	return len
}

func (LL *LinkedList) Delete(key interface{}){
	curr_node := LL.tail.next
	for curr_node != nil {
		if curr_node.next.value == key {
			curr_node.next = curr_node.next.next
			break
		}
		if curr_node == LL.tail {
			break
		}
		curr_node = curr_node.next
	}
}

func (LL *LinkedList) Reverse(){
	prev_node := LL.tail
	curr_node := LL.tail.next

	for LL.tail != nil {
		next_node := curr_node.next
		curr_node.next = prev_node
		prev_node = curr_node
		curr_node = next_node
		if curr_node == LL.tail.next {
			break
		}
		LL.tail = curr_node
	}
}


func (LL *LinkedList) ReverseRecursive(tail *Node) {
	curr_node := tail.next
	prev_node := tail
	if curr_node == LL.tail {
		return
	}
	LL.ReverseRecursive(curr_node)
	curr_node.next.next = curr_node
	curr_node.next = prev_node
	if prev_node == LL.tail {
		LL.tail = curr_node
	}

}

func (LL *LinkedList) Print(){
	temp := LL.tail.next
	for temp != nil {
		fmt.Print(temp.value)
		if temp.next == LL.tail.next{
			break
		}
		temp = temp.next
	}
}

func main() {
	ll := LinkedList{}
	ll.Append(6)
	ll.Append(4)
	ll.Append(9)
	ll.Append(4)
	ll.Append(9)
	ll.Print()
	////ll.Len()
	ll.Push(7)
	ll.Push(1)
	ll.Append(2)
	ll.Append(0)
	fmt.Println()
	ll.Print()
	fmt.Println(ll.tail)
	ll.Reverse()
	fmt.Println(ll.tail)
	fmt.Println()
	ll.Print()
	ll.ReverseRecursive(ll.tail)
	fmt.Println()
	ll.Print()
	fmt.Println()
	fmt.Println(ll.Len())
	ll.Delete(9)
	ll.Print()
	fmt.Println()
	fmt.Println(ll.Len())
	//ll.Append(4)
	//ll.Append(9)
	////ll.Insert(ll.head.next, 8)
	//ll.Print()
	//fmt.Println()
	//ll.Delete(1)
	//ll.Print()
	////ll.ReverseRecursive()
	//fmt.Println()
	//ll.Print()
	////ll.Reverse()
	//fmt.Println()
	//ll.Print()
	////fmt.Println(ll.head)
}