package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res []string
	var track []string
	// 变化的变量就是参数
	var generateParenthesisHelper func(int, int) //剩余的左和右括号数量
	generateParenthesisHelper = func(left int, right int) {
		if left < 0 || right < 0 {
			return
		}
		// 触发结束条件
		if left == 0 && right == 0 {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, strings.Join(tmp, ""))
			return
		}

		//for 选择 in 选择列表:
		//选择一个 "{"
		track = append(track, "{")
		//进入下一步决策
		generateParenthesisHelper(left-1, right)
		//撤销选择
		track = track[:len(track)-1]

		//选择一个 "}"
		track = append(track, "}")
		//进入下一步决策
		generateParenthesisHelper(left, right-1)
		//撤销选择
		track = track[:len(track)-1]
	}
	generateParenthesisHelper(n, n)
	return res
}
