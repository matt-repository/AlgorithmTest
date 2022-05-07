package main

import (
	"AlgorithmTest/LeedCode/general"
	"AlgorithmTest/LeedCode/simple"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	Test()
	//a:=[6]int{-1,0,3,5,9,11}
	//b:=[5]int{1,2,3,4,5}
	//c:=[1]int{1}
	//d:=[2]int{1,2}
	//fmt.Println(simple.Search1(a[:],2))
	//nums := []int{1,3,5,6,100}
	//target:= 101
	//fmt.Println(simple. SearchInsert(nums,target))

	//fmt.Println(simple.FirstBadVersion1(3))

	//result:=make([]int,0,100)
	//nums:=[]int{-4,-1,0,3,10}
	//simple.SortedSquares(nums)

	//fmt.Println(result,len(result),cap(result))
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//general.Rotate(nums[:], 3)
	//general.TwoSum([]int{5, 25, 75}, 100)
	//
	//simple.MaxSubArray([]int{8, -19, 5, -4, 20})
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//nums2 := []int{2, 5, 6}
	//simple.Merge(nums1, 3, nums2, 3)

	//simple.Intersect([]int{1,2,2,1},[]int{2,2})

	//simple.MaxProfit([]int{7,1,5,3,6,4})
	//
	//node:=simple.ListNode{Val: 1,
	//	Next: &simple.ListNode{Val:2,
	//		Next: &simple.ListNode{Val:3,
	//			Next: &simple.ListNode{Val:4,
	//				Next: &simple.ListNode{Val:5}}}}}
	// a:=node.Next.Next.Next.Next.Next
	//fmt.Println(a)
	//simple.MiddleNode(&node)
	// node:=&general.ListNode{Val: 1,
	//	 }
	// simple.ReverseList(&node)
	general.LengthOfLongestSubstring(
		"bbbbbb")
	simple.MatrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4)
	simple.Generate(5)

	general.CheckInclusion(
		"adc",
		"dcda")

	simple.FloodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	general.MaxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}})

	//simple.MergeTwoLists(&simple.ListNode{Val: 1,Next:&simple.ListNode{Val: 2,Next:&simple.ListNode{Val: 3}}},
	//	&simple.ListNode{Val:1,Next:&simple.ListNode{Val:3,Next:&simple.ListNode{Val:4}}})

	general.IsValidSudoku([][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '1', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '7'},
		{'.', '.', '.', '.', '.', '.', '.', '7', '.'}})

	general.SetZeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	})

	simple.CanConstruct("sss", "sss")
	general.LevelOrder(&general.TreeNode{Val: 1, Left: &general.TreeNode{Val: 2}, Right: nil})

	simple.IsSymmetric(&simple.TreeNode{Val: 1,
		Left: &simple.TreeNode{Val: 2,
			Right: &simple.TreeNode{Val: 3}},
		Right: &simple.TreeNode{Val: 2,
			Right: &simple.TreeNode{Val: 3}}})
	a := general.IsValidBST(&general.TreeNode{Val: 0})
	fmt.Println(a)
	//&general.TreeNode{Val: 3,Left:
	//&general.TreeNode{Val: 9},
	//	Right: &general.TreeNode{Val: 20,Left: &general.TreeNode{Val: 15},Right: &general.TreeNode{Val: 7}}}
	//fmt.Println( exchange([]int{1,2,3,4}))
	//data := jz_general.LevelOrder(&jz_general.TreeNode{
	//	Val: 3,
	//	Left: &jz_general.TreeNode{
	//		Val: 9,
	//	},
	//	Right: &jz_general.TreeNode{
	//		Val: 20,
	//		Left: &jz_general.TreeNode{
	//			Val: 15,
	//		},
	//		Right: &jz_general.TreeNode{
	//			Val: 7,
	//		},
	//	},
	//})
	//fmt.Println(data)
	fmt.Println((755204270 + 965550172) % 1000000007)
	//general.Combine(3,2)
	general.Permute([]int{1, 2, 3})
	//general.MinimumTotal([][]int{{-1},{2,3},{1,-1,-3}})

	fmt.Println(time.Now())

}

func Test() {
	obj := Constructor()
	obj.AddNum(6)
	obj.AddNum(10)
	obj.AddNum(2)
	obj.AddNum(6)
	obj.AddNum(5)
	obj.AddNum(0)
	obj.AddNum(6)
	obj.AddNum(3)
	obj.AddNum(1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())

	sb := &strings.Builder{}
	sb.WriteString(string(1))
	sb.WriteByte(',')
	sb.WriteString(strconv.Itoa(2))
	sb.WriteByte(',')
	sb.WriteString("null,")
	sb.WriteString(strconv.Itoa(1))
	sb.WriteByte(',')
	sb.WriteString(strconv.Itoa(2))
	sb.WriteByte(',')
	fmt.Println(sb.String())
	strconv.Atoi("")
}

type MedianFinder struct {
	Data []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		Data: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.Data) == 0 {
		this.Data = append(this.Data, num)
		return
	}
	start := 0
	last := len(this.Data) - 1
	for start < last {
		mid := (start + last) >> 1
		if this.Data[mid] > num {
			last = mid - 1
		} else if this.Data[mid] < num {
			start = mid + 1
		} else {
			start = mid
			break
		}
	}
	if num > this.Data[start] {
		start++
	}
	temp1 := make([]int, len(this.Data))
	copy(temp1, this.Data)
	temp2 := append(this.Data[:start], num)
	this.Data = append(temp2, temp1[start:]...)
}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.Data)
	if length%2 == 0 {
		return (float64)(this.Data[length/2]+this.Data[length/2-1]) / 2
	} else {
		return (float64)(this.Data[length/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
