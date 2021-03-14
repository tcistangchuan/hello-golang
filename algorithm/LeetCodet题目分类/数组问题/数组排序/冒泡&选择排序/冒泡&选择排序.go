package main

import "fmt"

/*	冒泡排序：
	冒泡算法(bubble sort) 是一种很简单的交换排序。每轮都从第一个元素开始，依次将较大值向后交换一位，直至整个数组有序。
*/
func bubbleSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums

}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	test := []int{12, 1, 3, 4, 2, 6, 5, 1}
	fmt.Println(selectSort(test))
}
