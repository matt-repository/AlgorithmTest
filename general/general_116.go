package general

//leedCode 116 题  填充每个节点的下一个右侧节点指针

///给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//初始状态下，所有next 指针都被设置为 NULL。

//输入：root = [1,2,3,4,5,6,7]
//输出：[1,#,2,3,#,4,5,6,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
//序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。

//输入：root = []
//输出：[]

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for queue != nil {

		tempQueue := []*Node{}
		for i := 0; i < len(queue); i++ {

			if i != len(queue)-1 {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tempQueue = append(tempQueue, queue[i].Right)
			}
		}

		queue = nil
		if len(tempQueue) > 0 {
			queue = tempQueue
		}
	}
	return root
}
