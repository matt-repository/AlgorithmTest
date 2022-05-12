package difficult

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := &strings.Builder{}
	var build func(root *TreeNode)
	build = func(root *TreeNode) {
		if root == nil {
			result.WriteString("nil,")
		} else {
			result.WriteString(strconv.Itoa(root.Val))
			result.WriteString(",")
			build(root.Left)
			build(root.Right)
		}

	}
	return result.String()

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	dataSource := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if len(dataSource) > 0 {
			v, ok := strconv.Atoi(dataSource[0])
			dataSource = dataSource[1:]
			if ok == nil {
				result := TreeNode{Val: v}
				result.Left = build()
				result.Right = build()
				return &result
			}
		}
		return nil
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
