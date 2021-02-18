package main
/*
int[] diff = new int[nums.length];
// 构造差分数组
diff[0] = nums[0];
for (int i = 1; i < nums.length; i++) {
	diff[i] = nums[i] - nums[i - 1];
}

int[] res = new int[diff.length];
// 根据差分数组构造结果数组
res[0] = diff[0];
for (int i = 1; i < diff.length; i++) {
	res[i] = res[i - 1] + diff[i];
}

原理很简单，回想diff数组反推nums数组的过程，diff[i] += 3意味着给nums[i..]所有的元素都加了 3
，然后diff[j+1] -= 3又意味着对于nums[j+1..]所有元素再减 3，那综合起来，是不是就是对nums[i..j]
中的所有元素都加 3 了？
*/
