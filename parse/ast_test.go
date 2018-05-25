package parse

import "testing"

func TestTree(t *testing.T) {
	tree := newTree()
	node := tree

	node.value = token{Value: "0"}

	node.left = newTree()
	node.left.value = token{Value: "l1"}
	node.right = newTree()
	node.right.value = token{Value: "r1"}

	node = node.left

	node.left = newTree()
	node.left.value = token{Value: "l2"}
	node.right = newTree()
	node.right.value = token{Value: "r2"}

	t.Logf("%v", tree)
}
