package main

import (
	"fmt"
)

func main(){
	//partition([]int{4,1,6,3,2,5},0,5,4)
	partition([]int{11,14,11,22,1,11},0,5)
	nums := []int{5,14,11,22,1,31,2,2,3,4,6}
	fmt.Println(sortArray(nums))

}


func sortArray(nums []int) []int {
	sort(nums,0,len(nums)-1)
	return nums
}

func sort(nums []int,left,right int) {
	if len(nums) == 0{
		return
	}
	if left >= right {
		return
	}
	p := partition(nums,left,right)
	sort(nums,left,p-1)
	sort(nums,p+1,right)
}

// 双指针
func partition(nums []int,left,right int)(p int){
	initVal := nums[left]
	initIndex := left
	for left!=right{
		for left<right && nums[right]>=initVal{
			right --
			fmt.Println("r",right)
		}
		nums[initIndex],nums[right] = nums[right],nums[initIndex]
		initIndex =right

		for left<right && nums[left]<initVal{
			left++
			fmt.Println("l",left)
		}
		nums[initIndex],nums[left] = nums[left],nums[initIndex]
		initIndex = left
	}
	fmt.Println(nums)
	return initIndex
}

