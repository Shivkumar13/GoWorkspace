package main

import "fmt"

// const n = 15
var nums []int

func main() {

	getNums()

	getConatenation(nums []int)

}

func getNums() {

	var n int

	fmt.Println("Enter the array size you want:")
	fmt.Scanln(&n)

	var nums = make([]int, n)

	fmt.Println("Now you can enter the values in array")

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
}

func getConatenation(nums []int) []int {

	n := len(nums)

	ans := make([]int, n*2)

	for i := 0; i < n; i++ {

		ans[i], ans[i+1] = nums[i], nums[i]

	}

	return ans
}
