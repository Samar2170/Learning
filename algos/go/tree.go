package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value, Left: nil, Right: nil}
	}
	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else if value > root.Value {
		root.Right = insert(root.Right, value)
	}
	return root
}
func (tree *BinaryTree) Insert(value int) {
	tree.Root = insert(tree.Root, value)
}

func inOrderTraversal(root *Node) {
	if root != nil {
		inOrderTraversal(root.Left)
		fmt.Printf("%d\n", root.Value)
		inOrderTraversal(root.Right)
	}
}

func (tree *BinaryTree) InOrderTraversal() {
	inOrderTraversal(tree.Root)
}

func preOrderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d\n", root.Value)
		preOrderTraversal(root.Left)
		preOrderTraversal(root.Right)
	}
}

func (tree *BinaryTree) PreOrderTraversal() {
	preOrderTraversal(tree.Root)
}

func postOrderTraversal(root *Node) {
	if root != nil {
		postOrderTraversal(root.Left)
		postOrderTraversal(root.Right)
		fmt.Printf("%d\n", root.Value)
	}
}

func (tree *BinaryTree) PostOrderTraversal() {
	postOrderTraversal(tree.Root)
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (tree *BinaryTree) Height() int {
	return height(tree.Root)
}

func search(root *Node, val int) bool {
	if root == nil {
		return false
	}
	if val < root.Value {
		return search(root.Left, val)
	}
	if val > root.Value {
		return search(root.Right, val)
	}
	return true
}

func (tree *BinaryTree) Search(val int) bool {
	return search(tree.Root, val)
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(20)
	tree.InOrderTraversal()

	fmt.Println("PreOrder Traversal")
	tree.PreOrderTraversal()

	fmt.Println("PostOrder Traversal")
	tree.PostOrderTraversal()

	fmt.Println("Height of the tree")
	fmt.Println(tree.Height())

	fmt.Println("Search for 7")
	fmt.Println(tree.Search(7))
	fmt.Println("Search for 17")
	fmt.Println(tree.Search(17))
}
