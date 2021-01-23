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
	if data == n.data {
		return
	}

	if data < n.data {
		if n.left != nil {
			n.left.Insert(data)
		}else {
			n.left = &node{data: data}
		}
	}else {
		if n.right != nil {
			n.right.Insert(data)
		}else {
			n.right = &node{data: data}
		}
	}
}

func (n *node) Search(value int) bool {
	if n.data == value {
		return true
	}
	if value < n.data {
		if n.left != nil {
			return n.left.Search(value)
		}else {
			return false
		}
	}
	if value > n.data {
		if n.right != nil {
			return n.right.Search(value)
		}else {
			return false
		}
	}
	return false
}

func (n *node) Max() int {
	if n.right != nil {
		return n.right.Max()
	}
	return n.data
}

func (n *node) Min() int {
	if n.left != nil {
		return n.left.Min()
	}
	return n.data
}

func (n *node) Delete(value int) *node {

	if value < n.data {
		if n.left != nil {
			n.left = n.left.Delete(value)
		}
	}else if value > n.data{
		if n.right != nil {
			n.right = n.right.Delete(value)
		}
	}else {
		if n.left == nil && n.right == nil {
			return nil
		}else if n.left == nil {
			n = n.right
			return n
		}else if n.right == nil {
			n = n.left
			return n
		}
		minRightSubTree := n.right.Min()
		n.data = minRightSubTree
		n.right = n.right.Delete(minRightSubTree)
	}

	return n
}


func (n *node) InOrderTraversal() []int {
	var res []int
	if n.left != nil {
		items := n.left.InOrderTraversal()
		res = append(res, items...)
	}

	res = append(res, n.data)

	if n.right != nil {
		items := n.right.InOrderTraversal()
		res = append(res, items...)
	}

	return res
}

func (n *node) PreOrderTraversal() []int {
	var res []int
	res = append(res, n.data)

	if n.left != nil {
		items := n.left.PreOrderTraversal()
		res = append(res, items...)
	}

	if n.right != nil {
		items := n.right.PreOrderTraversal()
		res = append(res, items...)
	}

	return res
}

func (n *node) PostOrderTraversal() []int {
	var res []int

	if n.left != nil {
		items := n.left.PostOrderTraversal()
		res = append(res, items...)
	}

	if n.right != nil {
		items := n.right.PostOrderTraversal()
		res = append(res, items...)
	}

	res = append(res, n.data)

	return res
}

func (n *node) InOrderTraversalRev() []int {
	var res []int

	if n.right != nil {
		items := n.right.InOrderTraversalRev()
		res = append(res, items...)
	}

	res = append(res, n.data)

	if n.left != nil {
		items := n.left.InOrderTraversalRev()
		res = append(res, items...)
	}

	return res
}

func main() {
	ip := []int{22,99,3,4,5,7,11,8,2,3,0,1,2,9}
	bt := BinaryTree(ip...)
	fmt.Println(bt.InOrderTraversal())
	bt.Insert(6)
	bt.Insert(10)
	fmt.Println(bt.InOrderTraversal())
	fmt.Println(bt.Search(23))
	fmt.Println(bt.Search(22))
	fmt.Println(bt.Max())
	fmt.Println(bt.Min())
	fmt.Println(bt.PreOrderTraversal())
	fmt.Println(bt.PostOrderTraversal())
	fmt.Println(bt.InOrderTraversal())
	fmt.Println(bt.Delete(22))
	fmt.Println(bt.InOrderTraversal())
	fmt.Println(bt.Delete(99))
	fmt.Println(bt.InOrderTraversal())

}