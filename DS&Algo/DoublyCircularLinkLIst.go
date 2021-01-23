package main

import "fmt"

type Node struct {
	prev *Node
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(key interface{}) {
	new_node := Node{
		prev: nil,
		data: key,
		next: nil,
	}

	if ll.head != nil {
		new_node.prev = ll.head.prev
		new_node.next = ll.head
		ll.head.prev.next = &new_node
		ll.head.prev = &new_node
	}else{
		ll.head = &new_node
		new_node.prev = ll.head
		new_node.next = ll.head
	}
}

func (ll *LinkedList) Push(key interface{}) {
	new_node := Node{
		prev: nil,
		data: key,
		next: nil,
	}

	if ll.head != nil {
		new_node.prev = ll.head.prev
		new_node.next = ll.head
		ll.head.prev.next = &new_node
		ll.head.prev = &new_node
		ll.head = &new_node
	}else{
		ll.head = &new_node
		new_node.prev = ll.head
		new_node.next = ll.head
	}
}

func (ll *LinkedList) Insert(index int, data interface{}) {
	len := ll.Len()
	if index <= len {
		var i int
		curr_node := ll.head
		if index == 0 {
			ll.Push(data)
			return
		}
		if index == len{
			ll.Append(data)
			return
		}
		for ll.head != nil {
			new_node := Node{
				prev: nil,
				data: data,
				next: nil,
			}
			i += 1
			if index == i {
				new_node.next = curr_node.next
				new_node.prev = curr_node
				curr_node.next.prev = & new_node
				curr_node.next = &new_node
				break
			}
			curr_node = curr_node.next

		}
	}
}

func (ll *LinkedList) Delete(index int){
	len := ll.Len()
	if index < len {
		curr_node := ll.head
		i := 0
		for ll.head != nil {
			if index == i {
				curr_node.next.prev = curr_node.prev
				curr_node.prev.next = curr_node.next
				if index == 0 {
					ll.head = curr_node.next
				}
				break
			}
			i += 1
			curr_node = curr_node.next
		}
	}
}

func (ll *LinkedList) Len() int{
	var len int
	head := ll.head
	for ll.head != nil {
		len += 1
		if head.next == ll.head {
			break
		}
		head = head.next
	}
	return len
}

func (ll *LinkedList) Reverse() {
	curr_node := ll.head
	for ll.head != nil {
		next_node := curr_node.next
		//next_node := curr_node.next
		curr_node.next = curr_node.prev
		curr_node.prev = next_node
		if curr_node.prev == ll.head{
			ll.head = curr_node
			break
		}
		curr_node = next_node
	}
}

func (ll *LinkedList) ReverseRecursive(head *Node) {
	curr_node := head
	if ll.head == nil || curr_node.next == ll.head{
		return
	}
	ll.ReverseRecursive(curr_node.next)
	curr_node.next.prev = curr_node.next.next
	curr_node.next.next = curr_node
	temp := curr_node.prev
	curr_node.prev = curr_node.next
	curr_node.next = temp
	if curr_node == ll.head{
		ll.head = curr_node.next
	}
}

func (ll *LinkedList) Print() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data)
		if temp == ll.head.prev{
			break
		}
		temp = temp.next
	}
}

func main() {
	ll := LinkedList{}
	ll.Append(9)
	ll.Append(8)
	ll.Append(7)
	ll.Append(6)
	ll.Print()
	ll.Push(10)
	fmt.Println()
	ll.Print()
	fmt.Println()
	fmt.Println(ll.Len())
	ll.Insert(5, 5)
	ll.Insert(0, 11)
	ll.Insert(3, "*")
	ll.Insert(0, "*")
	ll.Insert(9, "*")
	ll.Print()
	fmt.Println()
	ll.Delete(0)
	ll.Print()
	fmt.Println()
	ll.Delete(3)
	ll.Print()
	fmt.Println()
	ll.Delete(7)
	ll.Print()
	fmt.Println()
	ll.Reverse()
	ll.Print()
	fmt.Println(ll.head)
	ll.ReverseRecursive(ll.head)
	ll.Print()
	fmt.Println(ll.head)
}