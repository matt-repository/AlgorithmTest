package main

import (
	"AlgorithmTest/JianZhi_Offer/jz_difficult"
	"AlgorithmTest/JianZhi_Offer/jz_general"
	"AlgorithmTest/LeedCode/difficult"
	"AlgorithmTest/LeedCode/general"
	"AlgorithmTest/LeedCode/simple"
	"AlgorithmTest/Sort"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	jz_difficult.IsMatch("ab", "ab*")
	fmt.Println(Sort.Quick([]int{7, 5, 6, 4}))
	fmt.Println(jz_general.SingleNumber([]int{3, 4, 3, 3}))
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
	fmt.Println(difficult.SolveNQueens(1))
}

func Test() {
	i := 1
	i >>= 1
	fmt.Println(jz_general.NthUglyNumber(1))
	sb := &strings.Builder{}
	sb.WriteString(strconv.Itoa(1))
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
