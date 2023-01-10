package fib

func Tabulation(n int) int {
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

func Fibonacci(n int) int {
	return Tabulation(n)
}

func memo(n int, cache map[int]int) int {
	if n <= 1 {
		return n
	}

	if ans, ok := cache[n]; ok {
		return ans
	}

	ans := memo(n-1, cache) + memo(n-2, cache)

	cache[n] = ans

	return ans
}

func Memoization(n int) int {
	c := make(map[int]int)

	return memo(n, c)
}
