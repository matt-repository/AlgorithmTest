package jz_general

//复杂链表的复制

var mapCache = map[*Node]*Node{}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	tempNode, ok := mapCache[head]
	if ok {
		return tempNode
	}
	node := &Node{Val: head.Val}
	mapCache[head] = node
	node.Next = copyRandomList(head.Next)
	node.Random = copyRandomList(head.Random)

	return node
}

func replaceSpace(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result += "20%"
		} else {
			result += string(s[i])
		}
	}
	return result
}

func LevelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	q := []*TreeNode{root}
	p := []*TreeNode{}
	temp := []int{}
	for i := 0; i < len(q); i++ {
		if i == 0 {
			p = []*TreeNode{}
		}
		temp = append(temp, q[i].Val)

		if q[i].Left != nil {
			p = append(p, q[i].Left)
		}
		if q[i].Right != nil {
			p = append(p, q[i].Right)
		}
		if i == len(q)-1 {
			result = append(result, temp)

			q = p
			i = 0

		}
	}

	return result
}
