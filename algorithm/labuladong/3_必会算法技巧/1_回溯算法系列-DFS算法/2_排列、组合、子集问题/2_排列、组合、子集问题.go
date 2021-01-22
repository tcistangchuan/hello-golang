package main

import "fmt"

func main() {
	//fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(combine(4,4))
	fmt.Println(permute([]int{1, 2, 3}))

}

// 求子集-78题
func subsets(nums []int) [][]int {
	var res [][]int
	track := make([]int, 0)
	var subsetsHelper func(int)
	subsetsHelper = func(start int) {
		//触发结束条件
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		// 遍历选择列表
		for i := start; i < len(nums); i++ {
			// 排除不合法的选择
			// 做选择
			track = append(track, nums[i])
			// 进入下一层决策树 （注意这里是 i+1） ！！！！！！
			subsetsHelper(i + 1)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}
	subsetsHelper(0)
	return res
}

//求组合-77题
func combine(n int, k int) [][]int {
	var (
		nums  []int
		res   [][]int
		track []int
	)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var combineHelper func(int)
	combineHelper = func(index int) {
		if len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := index; i < len(nums); i++ {
			track = append(track, nums[i])
			combineHelper(i + 1)
			track = track[:len(track)-1]
		}

	}
	combineHelper(0)
	return res
}

// 全排列
func permute(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)
	var permuteHelper func()
	permuteHelper = func() {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		//fmt.Println("track:",track)
		for i := 0; i < len(nums); i++ {
			// 判断 是否合法
			if isExist(track,nums[i]){
				continue
			}
			track = append(track, nums[i])
			permuteHelper()
			track = track[:len(track)-1]
		}
	}
	permuteHelper()
	return res
}

// 判断是否包含数字
func isExist(nums []int,value int)bool{
	for _,v := range nums{
		if value == v{
			return true
		}
	}
	return false
}

