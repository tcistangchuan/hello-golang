package main

import (
	"fmt"
)

func main() {
	//partition([]int{4,1,6,3,2,5},0,5,4)
	//partition([]int{11, 14, 11, 22, 1, 11}, 0, 5)
	nums := []int{5, 14, 11, 22, 1, 31, 2, 2, 3, 4, 6}
	fmt.Println(sortArray(nums))

}

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

// 快速排序
func sort(nums []int, left, right int) {
	if len(nums) == 0 {
		return
	}
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	sort(nums, left, p-1)
	sort(nums, p+1, right)
}

// 快速选择排序
func quickSelectSort(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left < right {
		p := partition(nums, left, right)
		if k < p {
			right = p - 1
		} else if k > p {
			left = p + 1
		} else {
			return nums[p]
		}
	}
	return -1
}

// 双指针-快速排序
//快速排序的逻辑是，若要对nums[lo..hi]进行排序，我们先找一个分界点p，通过交换元素使得nums[lo..p-1]都小于等于nums[p]，
//且nums[p+1..hi]都大于nums[p]，然后递归地去nums[lo..p-1]和nums[p+1..hi]中寻找新的分界点，最后整个数组就被排序了
func partition(nums []int, left, right int) (p int) {
	initVal := nums[left]
	initIndex := left
	for left < right {
		for left < right && nums[right] >= initVal {
			right--
		}
		nums[initIndex], nums[right] = nums[right], nums[initIndex]
		initIndex = right

		for left < right && nums[left] < initVal {
			left++
		}
		nums[initIndex], nums[left] = nums[left], nums[initIndex]
		initIndex = left
	}
	fmt.Println(nums)
	return initIndex
}
