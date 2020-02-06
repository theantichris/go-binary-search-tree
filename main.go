package main

import (
	"errors"
	"fmt"
)

func main() {
	t := &TreeNode{value: 8}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	t.Delete(5)
	t.Delete(7)

	t.PrintInOrder()
	fmt.Println("")

	fmt.Printf("min is %d\n", t.FindMin())
	fmt.Printf("max is %d\n", t.FindMax())
}

// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Insert will insert a new value into the binary search tree.
func (treeNode *TreeNode) Insert(value int) error {
	if treeNode == nil {
		return errors.New("tree is empty")
	}

	if treeNode.value == value {
		return errors.New("this node value already exists")
	}

	if treeNode.value > value {
		if treeNode.left == nil {
			treeNode.left = &TreeNode{value: value}
			return nil
		}

		return treeNode.left.Insert(value)
	}

	if treeNode.value < value {
		if treeNode.right == nil {
			treeNode.right = &TreeNode{value: value}
			return nil
		}

		return treeNode.right.Insert(value)
	}

	return nil
}

// Find finds a value in the tree.
func (treeNode *TreeNode) Find(value int) (TreeNode, bool) {
	if treeNode == nil {
		return TreeNode{}, false
	}

	switch {
	case value == treeNode.value:
		return *treeNode, true
	case value < treeNode.value:
		return treeNode.left.Find(value)
	default:
		return treeNode.right.Find(value)
	}
}

// FindMin finds and returns the smallest value in the tree.
func (treeNode *TreeNode) FindMin() int {
	if treeNode.left == nil {
		return treeNode.value
	}

	return treeNode.left.FindMin()
}

// FindMax finds and returns the largest value in the tree.
func (treeNode *TreeNode) FindMax() int {
	if treeNode.right == nil {
		return treeNode.value
	}

	return treeNode.right.FindMax()
}

// Delete removes a value from the tree.
func (treeNode *TreeNode) Delete(value int) {
	treeNode.remove(value)
}

func (treeNode *TreeNode) remove(value int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if value < treeNode.value {
		treeNode.left = treeNode.left.remove(value)
		return treeNode
	}
	if value > treeNode.value {
		treeNode.right = treeNode.right.remove(value)
		return treeNode
	}

	if treeNode.left == nil && treeNode.right == nil {
		treeNode = nil
		return nil
	}

	if treeNode.left == nil {
		treeNode = treeNode.right
		return treeNode
	}
	if treeNode.right == nil {
		treeNode = treeNode.left
		return treeNode
	}

	smallestValOnRight := treeNode.right

	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	treeNode.value = smallestValOnRight.value
	treeNode.right = treeNode.right.remove(treeNode.value)

	return treeNode
}

// PrintInOrder transverses the tree and prints the values in order.
func (treeNode *TreeNode) PrintInOrder() {
	if treeNode == nil {
		return
	}

	treeNode.left.PrintInOrder()
	fmt.Println(treeNode.value)
	treeNode.right.PrintInOrder()
}
