package math

/*
961. 在长度 2N 的数组中找出重复 N 次的元素：https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/

给你一个整数数组 nums ，该数组具有以下属性：
- nums.length == 2 * n.
- nums 包含 n + 1 个 不同的 元素
- nums 中恰有一个元素重复 n 次

找出并返回重复了 n 次的那个元素。

示例 1：
	输入：nums = [1,2,3,3]
	输出：3

解题思路：
	1. 哈希表：使用一个 set 存储已经出现过的元素
	2. 摩尔投票：从 nums[1, n-1] 中寻找多数元素
*/

// 摩尔投票
//   - 若 nums[0] 在 nums[1, n-1] 中重复出现，则 nums[0] 为所需元素
//   - 否则在 nums[1, n-1] 中寻找多数元素
func RepeatedNTimes(nums []int) int {
	candidate, count := 0, 0
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		// 重复出现，nums[0] 为所需元素
		if v == nums[0] {
			return v
		}
		if count == 0 {
			candidate, count = v, 1
			continue
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}
	return candidate
}
