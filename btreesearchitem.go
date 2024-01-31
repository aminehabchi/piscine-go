package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if root.Left != nil {
		if elem != root.Left.Data {
			BTreeSearchItem(root.Left, elem)
		} else {
			return root.Left
		}
	}
	if root.Right != nil {
		if elem != root.Right.Data {
			BTreeSearchItem(root.Right, elem)
		} else {
			return root.Right
		}
	}
	return root
}
