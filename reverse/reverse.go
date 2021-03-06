package main

import "fmt"

func main() {
	a := reverse2(123)
	fmt.Println(a)
}

func reverse(x int) int {
	var nums []int
	var ans int
	tmp := x
	if x < 0 {
		x = -x
	}
	for x > 0 {
		remainder := x % 10
		nums = append(nums, remainder)
		x = x / 10
	}
	i := 0
	for n := len(nums) - 1; n >= 0; n-- {
		ans += nums[i] * pow(n)
		i++
	}
	if tmp < 0 {
		ans = -ans
	}
	if ans > 2147483647 || ans < -2147483648 {
		ans = 0
	}
	return ans
}

func pow(n int) int {
	if n == 0 {
		return 1
	} else {
		return 10 * pow(n-1)
	}
}

func reverse2(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > 2147483647 || res < -2147483648 {
		res = 0
	}
	return res
}
