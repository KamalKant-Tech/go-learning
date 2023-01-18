package main

import "fmt"

func main() {

	// taking an array variable
	// var n [100]float64
	// var total int
	// fmt.Print("Enter number of elements: ")

	// // taking user input
	// fmt.Scanln(&total)
	// for i := 0; i < total; i++ {
	// 	fmt.Print("Enter the number : ")
	// 	fmt.Scan(&n[i])
	// }
	// for j := 1; j < total; j++ {
	// 	if n[0] < n[j] {
	// 		n[0] = n[j]
	// 	}

	// }

	// fmt.Print("The largest number is : ", n[0])
	// sliceInt := []int{5, 4, -1, 7, 8}
	// sliceInt := []int{5, 4, -1, 7, 8}
	sliceInt := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceInt)
	fmt.Println(InsertionSort(sliceInt))
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	storeNumbers := map[int]int{}
	for i := 0; i < len(nums); i++ {
		storeNumbers[nums[i]] += 1
	}
	fmt.Println(storeNumbers)
	for _, v := range storeNumbers {
		if v > 1 {
			return true
		}
	}
	return false
}

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// A subarray is a contiguous part of an array.
func maxSubArray(nums []int) int {
	max := nums[0]
	min := nums[0]

	for i := 0; i < len(nums); i++ {
		min = min + nums[i]

		if max < min {
			max = min
		}
		if min < 0 {
			min = 0
		}
	}
	return max
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+nums[i+1] == target {
			return []int{i, i + 1}
		}
	}
	return []int{0, 0}
}

// Bubble sort
func sortArray(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				fmt.Println(nums[j], nums[j-1])
				tmp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = tmp
			}
		}
	}
	return nums
}

// Insertion sort: pick an lowest or largest element and put this into their respective index
func InsertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		lastIndex := len(nums) - i - 1
		maxIndex := getMaxEleIndex(nums, 0, lastIndex)
		fmt.Println(lastIndex, maxIndex)
		tmp := nums[lastIndex]
		nums[lastIndex] = nums[maxIndex]
		nums[maxIndex] = tmp
	}
	return nums
}

func getMaxEleIndex(nums []int, start int, last int) int {
	max := start

	for i := 0; i <= last; i++ {
		if nums[max] < nums[i] {
			max = i
		}
	}
	return max
}
