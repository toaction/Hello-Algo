package math

/*
169. 多数元素：https://leetcode.cn/problems/majority-element/

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：
	输入：nums = [3,2,3]
	输出：3

解题思路：
	1. 摩尔投票算法
*/

// 摩尔投票算法：
//   - 把多数元素和其他元素“配对抵消”
//   - 因为多数元素数量 > 其他所有元素数量之和，所以抵消后剩下的一定是多数元素
func MajorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count = 1
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
