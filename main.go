package main

import "errors"

func main() {

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

// FindMin finds and returns the smallets value in the tree.
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
