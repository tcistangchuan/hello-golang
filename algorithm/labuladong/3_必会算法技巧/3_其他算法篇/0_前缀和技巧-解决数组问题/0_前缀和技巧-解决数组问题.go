package main

import (
	"fmt"
)

func main(){
	fmt.Println(subarraySum([]int{1,1,1},2))
	fmt.Println(subarraySum2([]int{1,1,1},2))
	//fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

// 求连续子数组 满足k的个数 时间复杂度O(N^2)
func subarraySum(nums []int, k int) int {
	total := 0
	preNums := make([]int,len(nums)+1)
	preNums[0] = 0
	for i:=0;i<len(nums);i++{
		preNums[i+1] =nums[i]+preNums[i]
	}
	for j:=1;j<=len(nums);j++ {
		for i := 0; i < j; i++ {
			// 判断 [i:j-1] 是否满足 k
			if preNums[j]-preNums[i] == k{
				total++
			}
		}
	}
	return total
}

// 借助 map 优化
func subarraySum2(nums []int, k int) int {
	total := 0
	preNums := make([]int,len(nums)+1)
	preNums[0] = 0
	preNumsMap := make(map[int]bool)
	preNumsMap[0] = true
	for i:=0;i<len(nums);i++{
		preNums[i+1] =nums[i]+preNums[i]
		preNumsMap[preNums[i+1]] = true
	}
	sum_j := 0
	for j:=0;j<len(nums);j++{
		sum_j += nums[j]
		if preNumsMap[sum_j-k] {
			total++
		}
	}
	return total
}


