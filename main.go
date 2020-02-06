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
