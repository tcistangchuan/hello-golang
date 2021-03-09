package q1_两数之和

//1. q1_两数之和（数组是无序的）
//一、暴力破解法  时间复杂度是 O(n^2)
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//二、借用map辅助方法
/*
1.构造一个哈希表。将数组放入map中 map[value]key
2.遍历数组判断这样的map中的元素是否存在:
	// 元素key等于target-nums[i]存在 且不等于 本身nums[i]
	if map[target-nums[i]] > 0  && m[target-nums[i]]!=i{
		return []int{i,map[target-nums[i]]}
	}
*/
func twoSum2(nums []int, target int) []int {
	n := len(nums)
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	for i := 0; i < n; i++ {
		if m[target-nums[i]] > 0  && m[target-nums[i]]!=i{
			return []int{i, m[target-nums[i]]}
		}
	}
	return []int{}
}


// 扩展：若数组是有序的
/*
 双指针解法：
int[] twoSum(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left < right) {
        int sum = nums[left] + nums[right];
        if (sum == target) {
            return new int[]{left, right};
        } else if (sum < target) {
            left++; // 让 sum 大一点
        } else if (sum > target) {
            right--; // 让 sum 小一点
        }
    }
    // 不存在这样两个数
    return new int[]{-1, -1};
}

*/
