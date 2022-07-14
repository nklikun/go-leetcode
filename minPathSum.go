package main

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例 1：
1 3 1
1 5 1
4 2 1

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：
1 2 3
4 5 6
输入：grid = [[1,2,3],[4,5,6]]
输出：12


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run64() {}

func minPathSum(grid [][]int) int {
	// make a dp, but not use gird ^.^
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				// do nothing
			} else if i == 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 {
				grid[i][0] += grid[i-1][0]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
