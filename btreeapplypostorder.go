package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if root.Left != nil {
		if root.Left.Data != "" {
			f(root.Left.Data)
			root.Left.Data = ""
			BTreeApplyPostorder(root.Left, f)
		}
	}
	if root.Right != nil {
		if root.Right.Data != "" {
			f(root.Right.Data)
			root.Right.Data = ""
			BTreeApplyPostorder(root.Right, f)
		}
	}
	if root.Data != "" {
		f(root.Data)
		root.Data = ""
	}
}
