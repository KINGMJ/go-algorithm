package climbingstairs

// climbStairs 计算爬n阶楼梯的不同方法数
// n: 楼梯阶数
// return: 不同的爬楼梯方法数

// 递归解法：其实就是斐波那契数列的问题
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 递归，带备忘录解法
func climbStairs2(n int) int {
	memo := make(map[int]int)
	return climb(n, memo)
}

func climb(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	if res, ok := memo[n]; ok {
		return res
	}
	memo[n] = climb(n-1, memo) + climb(n-2, memo)
	return memo[n]
}

// 动态规划，状态转移方程
// dp[i] = dp[i-1] + dp[i-2]
// dp[1] = 1
// dp[2] = 2
func climbStairs3(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
