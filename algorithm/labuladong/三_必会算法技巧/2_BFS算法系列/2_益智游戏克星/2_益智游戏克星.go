package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
	fmt.Println(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
}

func slidingPuzzle(board [][]int) int {
	var q [][][]int
	end := [][]int{{1, 2, 3}, {4, 5, 0}}
	q = append(q, board)
	step := 0
	//不回头
	visited := make(map[string]bool, 0)
	visited[slice2str(board)] = true
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			// 是否到达终点
			if reflect.DeepEqual(cur, end) {
				return step
			}
			// 检查是否需要继续执行
			// 添加相邻节点
			for _, v := range getNodes(cur) {
				if !visited[slice2str(v)] {
					q = append(q, v)
					visited[slice2str(v)] = true
				}
			}
		}
		step++
	}
	return -1
}

func slice2str(nums [][]int) string {
	s := ""
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			s += strconv.Itoa(nums[i][j])
		}
	}
	return s
}

// 找出所有相邻节点
func getNodes(nums [][]int) (res [][][]int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == 0 {
				if j-1 >= 0 {
					res = append(res, change(nums, i, j, i, j-1))
				}
				if j+1 < len(nums[i]) {
					res = append(res, change(nums, i, j, i, j+1))
				}
				if i-1 >= 0 {
					res = append(res, change(nums, i, j, i-1, j))
				}
				if i+1 < len(nums) {
					res = append(res, change(nums, i, j, i+1, j))
				}
			}
		}
	}
	return res
}

func change(nums [][]int, i, j, i2, j2 int) [][]int {
	tmp := make([][]int, len(nums))
	copy(tmp, nums)
	for i := 0; i < len(nums); i++ {
		tmp2 := make([]int, len(nums[i]))
		copy(tmp2, nums[i])
		tmp[i] = tmp2
	}
	tmp[i][j], tmp[i2][j2] = tmp[i2][j2], tmp[i][j]
	return tmp
}
