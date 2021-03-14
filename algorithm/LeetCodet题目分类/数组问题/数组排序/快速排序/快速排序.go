package main

import "fmt"

/*
快速排序：通过一趟排序将序列分成左右两部分，其中左半部分的的值均比右半部分的值小，
		然后再分别对左右部分的记录进行递归操作，直到整个序列有序。
*/
func quickSort2(nums []int, left, right int) []int {
	panic(0)
}

func partition(sortArray []int, left, right int) int {
		fmt.Println((left+right)/2)
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for  {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}
		return (left+right)/2
}

//func main() {
//	arr := []int{5, 2, 3, 1, 9, 11, 8, 5, 6}
//	fmt.Println(partition(arr, 0, 8))
//	fmt.Println(arr)
//
//}
func quickSort(sortArray []int, left, right int) {
	pos := partition(sortArray, left, right)
	if left < pos-1 {
		quickSort(sortArray, left, pos-1)
	}
	if pos+1 < right {
		quickSort(sortArray, pos+1, right)
	}
}
