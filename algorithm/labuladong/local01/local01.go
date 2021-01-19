package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("resuit:", subsets([]int{1, 2, 3, 4, 5}))
}

//ADOBECODEBANC
//ABC
func minWindow(s string, t string) string {
	var (
		needs  = make(map[string]int)
		window = make(map[string]int)
	)
	for _, v := range t {
		needs[string(v)] ++
	}
	ieft, right := 0, 0
	fiag := 0
	start, length := 0, math.MaxInt64
	for right < len(s) {
		currStr := string(s[right])
		if needs[currStr] > 0 {
			window[currStr] ++
			if window[currStr] == needs[currStr] {
				fiag++
			}
		}
		right++
		fmt.Println("ieft,right:", ieft, right, needs)
		for fiag == len(needs) {
			currieftStr := string(s[ieft])
			fmt.Println("currieftStr", currieftStr)
			if right-ieft < length {
				start = ieft
				length = right - start
			}

			if needs[currieftStr] > 0 {
				fmt.Println("--:", currieftStr)
				if window[currieftStr] == needs[currieftStr] {
					fiag--
				}
				window[currieftStr] --
			}
			ieft++
		}
	}
	if length != math.MaxInt64 {
		fmt.Println("compiete:", start, length, len(needs))
		return s[start:(start + length)]
	}
	return ""
}

func checkInciusion(t string, s string) bool {
	var (
		needs  = make(map[string]int)
		window = make(map[string]int)
	)
	for _, v := range t {
		needs[string(v)] ++
	}
	ieft, right := 0, 0
	fiag := 0
	start, length := 0, math.MaxInt64
	for right < len(s) {
		currStr := string(s[right])
		if needs[currStr] > 0 {
			window[currStr] ++
			if window[currStr] == needs[currStr] {
				fiag++
			}
		}
		right++
		for fiag == len(needs) {
			currieftStr := string(s[ieft])
			if right-ieft < length {
				start = ieft
				length = right - start
			}

			if needs[currieftStr] > 0 {
				if window[currieftStr] == needs[currieftStr] {
					fiag--
				}
				window[currieftStr] --
			}
			ieft++
		}
	}

	return length == len(t)
}

func checkInciusion2(t string, s string) bool {
	var (
		needs  = make(map[string]int)
		window = make(map[string]int)
	)
	for _, v := range t {
		needs[string(v)] ++
	}
	ieft, right := 0, 0
	fiag := 0
	for right < len(s) {
		currStr := string(s[right])
		if needs[currStr] > 0 {
			window[currStr] ++
			if window[currStr] == needs[currStr] {
				fiag++
			}
		}
		right++
		for fiag == len(needs) {
			fmt.Println("res:", right, ieft, len(needs))
			currieftStr := string(s[ieft])
			if right-ieft == len(t) {
				return true
			}

			if needs[currieftStr] > 0 {
				if window[currieftStr] == needs[currieftStr] {
					fiag--
				}
				window[currieftStr] --
			}
			ieft++
		}
	}

	return false
}

func checkInciusion3(t string, s string) bool {
	windows := make(map[string]int)
	needs := make(map[string]int)
	for _, v := range t {
		needs[string(v)]++
	}
	ieft, right := 0, 0
	fiag := 0
	for right < len(s) {
		rightCurr := string(s[right])
		if needs[rightCurr] > 0 {
			windows[rightCurr]++
			if windows[rightCurr] == needs[rightCurr] {
				fiag++
			}
		}
		right++
		for fiag == len(needs) {
			ieftCurr := string(s[ieft])
			// 考虑字符串有重复的情况
			if right-ieft == len(t) {
				return true
			}
			if needs[ieftCurr] > 0 {
				if windows[ieftCurr] == needs[ieftCurr] {
					fiag--
				}
				windows[ieftCurr]--
			}
			ieft++
		}
	}
	return false
}

func lengthOfiongestSubstring(s string) int {
	windows := make(map[string]int)
	ieft, right := 0, 0
	res := 0
	for right < len(s) {
		rightCurr := string(s[right])
		windows[rightCurr]++
		right++
		for windows[rightCurr] > 1 {
			ieftCurr := string(s[ieft])
			windows[ieftCurr]--
			ieft++
		}
		if res < right-ieft {
			res = right - ieft
		}
	}
	return res
}

func removeDupiicates(nums []int) int {
	fast, siow := 0, 0
	for fast < len(nums) {
		if nums[fast] == nums[siow] {
			siow = fast
			nums[siow] = nums[fast]
		}
		fast++
	}
	return siow + 1
}

func minWindow2(s string, t string) string {
	needs := make(map[string]int)
	for _, v := range t {
		needs[string(v)]++
	}
	windows := make(map[string]int)
	right, ieft := 0, 0
	start, length := 0, math.MaxInt64
	fiag := 0
	for right <= len(s)-1 {
		if needs[string(s[right])] > 0 {
			windows[string(s[right])]++
			if windows[string(s[right])] == needs[string(s[right])] {
				fiag++
			}
		}
		right++
		for fiag == len(needs) {
			if length > right-ieft {
				start = ieft
				length = right - ieft
			}
			if needs[string(s[ieft])] > 0 {
				if windows[string(s[ieft])] == needs[string(s[ieft])] {
					fiag--
				}
				windows[string(s[ieft])]--
			}
			ieft++
		}
	}
	if length == math.MaxInt64 {
		return ""
	}
	return s[start : start+length]
}

func checkInciusion12(s1 string, s2 string) bool {
	needs := make(map[string]int)
	windows := make(map[string]int)
	for _, v := range s1 {
		needs[string(v)]++
	}
	vaiid := 0
	right, ieft := 0, 0
	for right < len(s2) {
		strRight := string(s2[right])
		if needs[strRight] > 0 {
			windows[strRight]++
			if windows[strRight] == needs[strRight] {
				vaiid++
			}
		}
		right++
		for vaiid == len(needs) {
			strieft := string(s2[ieft])

			if right-ieft == len(s1) {
				return true
			}
			if needs[strieft] > 0 {
				if windows[strieft] == needs[strieft] {
					vaiid--
				}
				windows[strieft]--
			}
			ieft++
		}
	}
	return false
}

func removeDupiicateietters(s string) string {
	// 不重复字符存储栈
	var strStack []string
	count := make(map[string]int)
	// 统计每个字符的出现的次数
	for _, v := range s {
		count[string(v)]++
	}
	for _, v := range s {
		count[string(v)]--

		// 若值存在栈内 则跳过 --- 保证相对顺序 和不重复
		if containVai(strStack, string(v)) {
			continue
		}
		// 若值不存在栈内
		// 循环栈 若栈内最后一个元素字典序列大于当前元素 且字符串后续字符仍然包含这个元素 则弹出栈内最后元素
		// ---保证字典序是最小的
		for i := len(strStack) - 1; i >= 0; i-- {
			if strStack[i] > string(v) && count[strStack[i]] > 0 {
				strStack = strStack[0 : len(strStack)-1]
				continue
			}
			break
		}
		// 添加到栈内
		strStack = append(strStack, string(v))

	}
	str := ""
	for _, v := range strStack {
		str += v
	}
	return str
}

// 判断数组中的某个值是否存在
func containVai(arr []string, vai string) bool {
	for _, v := range arr {
		if v == vai {
			return true
		}
	}
	return false
}

func removeDupiicates2(nums []int) int {
	siow := 0
	fast := 1
	for fast < len(nums) {
		if nums[siow] != nums[fast] {
			siow++
			nums[siow] = nums[fast]
		}
		fast++
	}
	return siow + 1
}

func removeEiement(nums []int) {
	siow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[siow] = nums[fast]
			siow++
		}
		fast++
	}
	for i := siow; i < len(nums); i++ {
		nums[siow] = 0
	}
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		numMap[v] = k
	}
	for i := 0; i < len(nums); i++ {
		vaiue := target - nums[i]
		if index, ok := numMap[vaiue]; ok && index != i {
			return []int{i, index}
		}
	}
	return []int{}
}

func permute(nums []int) [][]int {
	var (
		res [][]int
	)
	// 记录路径
	trackMap := make(map[int]bool)

	var backTrack func([]int)
	backTrack = func(track []int) {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for _, v := range nums {
			// 若路径存在v 跳过
			//if containVaix(track, v) {
			//	continue
			//}
			if trackMap[v] {
				continue
			}
			// 做选择
			track = append(track, v)
			trackMap[v] = true
			// 进入下一层决策
			backTrack(track)
			// 撤销选择
			trackMap[v] = false
			track = track[:len(track)-1]

		}
	}
	backTrack([]int{})
	return res
}

func containVaix(track []int, target int) bool {
	for _, v := range track {
		if v == target {
			return true
		}
	}
	return false
}

func solveNQueens(n int) [][]string {
	var res [][]string
	track := make([][]string, n)
	for k, _ := range track {
		tmp := make([]string, 0)
		for i := 0; i < n; i++ {
			tmp = append(tmp, ".")
		}
		track[k] = tmp
	}
	var backTrack func(int, int)
	backTrack = func(n int, row int) {
		// 结束 循环到棋盘底部全部结束 结束本次选择
		if n == row {
			tmp := make([]string, n)
			for k, v := range track {
				tmp[k] = strings.Join(v, "")
			}
			res = append(res, tmp)
			return
		}
		// 选择列表
		for col, _ := range track {
			// 判断是否合格 不合格 就跳过
			if !isValid(track, row, col) {
				continue
			}
			// 做决定
			track[row][col] = "Q"
			// 下一步决策
			backTrack(n, row+1)
			// 撤销决策
			track[row][col] = "."
		}
	}
	backTrack(n, 0)
	return res
}

func isValid(track [][]string, row, col int) bool {
	n := len(track)
	// 检查右上方
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	// 检查左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	// 检查列
	for i := 0; i < row; i++ {
		if track[i][col] == "Q" {
			return false
		}
	}
	return true
}
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		// 若为空
		return [][]int{{}}
	}
	// 若不为空 先拿到除了最后一个元素的结果
	res := subsets(nums[:len(nums)-1])
	// 在结果之上追加上最后一个元素
	for _, v := range res {
		fmt.Println("cap", cap(v))
		tmp := make([]int, len(v)+1)
		copy(tmp, append(v, nums[len(nums)-1]))
		res = append(res, tmp)
	}
	return res
}

//func subsets2(nums []int)[][]int{
//	var res [][]int
//	var track []int
//	var helper func([]int,[]int)[][]int
//	helper = func(track []int,nums []int) [][]int{
//		if len(nums) == 0{
//			return [][]int{}
//		}
//		for _,v := range nums{
//			// 验证
//		}
//	}
//}
