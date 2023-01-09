package fib

func Fibonacci(n int) int {
	if n == 0 {
		return n
	}

	dp := make([]int, n+1)

	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
