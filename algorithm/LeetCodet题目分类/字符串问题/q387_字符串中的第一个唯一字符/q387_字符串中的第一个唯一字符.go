package q387_字符串中的第一个唯一字符
/*
方法一、使用哈希表存储频数
思路与算法
	我们可以对字符串进行两次遍历。
	在第一次遍历时，我们使用哈希映射统计出字符串中每个字符出现的次数。
	在第二次遍历时，我们只要遍历到了一个只出现一次的字符，那么就返回它的索引，否则在遍历结束后返回 -1。
时间复杂度：O(n)

*/

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _,v := range s{
		m[v]++
	}
	for i,v :=range s{
		if m[v] == 1{
			return i
		}
	}
	return -1
}