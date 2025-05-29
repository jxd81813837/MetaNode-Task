package Goroutine

import "fmt"

func main() {
	go shouOdd([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	go shouEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

}

func shouOdd(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Println("Odd:", nums[i])
		}
	}
}
func shouEven(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			fmt.Println("Even:", nums[i])
		}
	}
}
