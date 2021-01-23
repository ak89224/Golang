package main

import (
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func BinaryTree(items ...int) *node {
	root := node{data: items[0]}
	if len(items) > 1 {
		for _, data := range items[1:] {
			root.Insert(data)
		}
	}
	return &root
}

func (n *node) Insert(data int)  {
	if n != nil {
		var queue []*node
		queue = append(queue, n)

		for len(queue) != 0 {
			root := queue[0]
			queue = queue[1:]

			if root.left != nil {
				queue = append(queue, root.left)
			}else {
				root.left = BinaryTree(data)
				return
			}

			if root.right != nil {
				queue = append(queue, root.right)
			}else {
				root.right = BinaryTree(data)
				return
			}
		}
	}
}

func (n *node) Inorder() {
	if n != nil {
		n.left.Inorder()
		fmt.Print(n.data, "  ")
		n.right.Inorder()
	}else {
		return
	}
}

func (n *node) LevelOrder() {
	if n != nil {
		var queue []*node
		queue = append(queue, n)

		for len(queue) != 0 {
			root := queue[0]
			queue = queue[1:]

			fmt.Print(root.data, "  ")

			if root.left != nil {
				queue = append(queue, root.left)
			}
			if root.right != nil {
				queue = append(queue, root.right)
			}
		}
	}else {
		return
	}
}

func (n *node) Delete(key int) {
	if n != nil {
		//only 1 item in the tree.
		if n.left == nil && n.right == nil {
			if n.data == key {
				n = nil
			}
			return
		}

		//level-order to find the last node
		var queue []*node
		var key_node *node
		var last_node *int
		queue = append(queue, n)

		for len(queue) != 0 {
			root := queue[0]
			queue = queue[1:]

			if root.data == key{
				key_node = root
			}

			if root.right != nil {
				if root.right.right == nil && root.right.left == nil && last_node == nil{
					last_node = &root.right.data
					root.right = nil
				}else {
					queue = append(queue, root.right)
				}
			}

			if root.left != nil {
				if root.left.right == nil && root.left.left == nil && last_node == nil{
					last_node = &root.left.data
					root.left = nil
				}else {
					queue = append(queue, root.left)
				}
			}

			if key_node != nil && last_node != nil {
				key_node.data = *last_node
				break
			}
		}
	}else {
		return
	}
}

func (n *node) GetLevel(key int) int {
	level := 0
	if n != nil{
		level = 1
		if n.data == key{
			return level
		}

		level += n.left.GetLevel(key)
		if level > 1 {
			return level
		}

		level += n.right.GetLevel(key)
		if level > 1 {
			return level
		}

		return 0
	}else {
		return level
	}
}

func (n *node) GetHeight() int {
	heightL := 0
	heightR := 0
	if n != nil {
		heightL = n.left.GetHeight()

		heightR = n.right.GetHeight()

		//Save the height of the current node before each return if you want.

		if heightL > heightR {
			return heightL + 1
		}else {
			return heightR + 1
		}

		//For size or no. of nodes
		//size = heightL + 1 + heightR
	}else {
		return 0
	}
}

func main() {
	bt := BinaryTree(1,1,2,3,4,5,5)
	bt.Inorder()
	fmt.Println()
	bt.Insert(6)
	bt.Inorder()
	fmt.Println()
	bt.Insert(9)
	bt.Inorder()
	fmt.Println()
	bt.LevelOrder()
	fmt.Println()
	bt.Delete(5)
	bt.LevelOrder()
	fmt.Println()
	fmt.Println(bt.GetLevel(5))
	fmt.Println(bt.GetLevel(1))
	fmt.Println(bt.GetLevel(2))
	fmt.Println(bt.GetLevel(3))
	fmt.Println(bt.GetLevel(4))
	fmt.Println(bt.GetLevel(6))
	fmt.Println(bt.GetLevel(9))
	fmt.Println(bt.GetHeight(), "<--")



}