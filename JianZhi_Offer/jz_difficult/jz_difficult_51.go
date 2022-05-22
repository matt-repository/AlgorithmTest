package jz_difficult

//数组中的逆序对


func ReversePairs(nums []int) int {
	if len(nums)<2{
		return 0
	}
	temp:=make([]int,len(nums))
	copy(temp,nums)
	count,_:=mergeSort(nums,0,len(nums)-1,temp)
	return count

}

func mergeSort(nums []int, start,end int ,temp []int)(int,[]int) {
	if start==end{
		return 0,nums
	}
	mid:=start+(end-start)/2
	leftCount,tempNums1:=mergeSort(nums,start,mid,temp )
	rightCount,tempNums2:=mergeSort(tempNums1,mid+1,end,temp )

	if tempNums2[mid]<=tempNums2[mid+1] {
		return leftCount+rightCount,tempNums2
	}
	crossPairs,tempNums3:=mergeCount(tempNums2,start,mid,end,temp)
	return leftCount+rightCount+crossPairs,tempNums3

}

func  mergeCount(nums []int, start,mid,end int ,temp []int) (int,[]int) {
	for i:=start;i<=end;i++{
		temp[i]=nums[i]
	}

	i:=start
	j:=mid+1
	count:=0
	for k:=start;k<=end;k++{
		if i==mid+1{
			nums[k]=temp[j]
			j++
		}else if j==end+1{
			nums[k]=temp[i]
			i++
		}else if temp[i]<=temp[j]{
			nums[k]=temp[i]
			i++
		}else {
			nums[k]=temp[j]
			j++
			count+=(mid-i+1)
		}
	}
	return count,nums
}


