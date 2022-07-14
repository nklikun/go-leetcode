package main

import "fmt"

// https://leetcode.cn/problems/jump-game-ii/
/*
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

1 <= nums.length <= 104
0 <= nums[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/jump-game-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run45() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = i + 1
	}
	for i := 0; i < len(nums); i++ {
		// 从i最远可以跳到i+nums[i]， 那么i-i+nums[i]这段全部可以从i一步跳到
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			if dp[i]+1 < dp[j] {
				dp[j] = dp[i] + 1
			}
		}
	}
	return dp[len(nums)-1]
}
