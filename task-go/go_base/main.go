package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(towSum([]int{2, 7, 11, 12, 15}, 23))
}

func towSum(nums []int, target int) []int {
	sort.Ints(nums) //对数组进行排序
	for i := len(nums) / 2; i > 0 || i < len(nums); {
		if nums[i]+nums[i-1] > target {
			i /= 2
		} else if nums[i]+nums[i-1] < target {
			i = i + i/2
		} else if nums[i]+nums[i-1] == target {
			return []int{nums[i], nums[i-1]}
		} else if nums[i]+nums[i+1] == target {
			return []int{nums[i], nums[i+1]}
		}
	}
	return []int{}
}

func mergeIntervals(intervals [][]int) [][]int {
	intervalsResult := [][]int{}
	for i := 0; i < len(intervals)-1; i++ {
		first := intervals[i][1]  //取出前一个元素最后一个
		last := intervals[i+1][0] //取出后一个元素第一个
		if first >= last {
			arr := []int{intervals[i][0], intervals[i+1][1]}
			intervalsResult = append(intervalsResult, arr)
		}
	}
	return intervalsResult
}

func plusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
		if i == 0 {
			return append([]int{1}, nums...)
		}
	}
	return nums
}

// 输出 删除排序数组中的重复项
func removeDuplicates(nums []int) (int, []int) {
	resultInt := []int{}
	resultCount := 0
	resultInt = append(resultInt, nums[0]) //把第一个元素追加进去
	for i := 1; i < len(nums); i++ {
		countTemp := nums[i]
		for j := resultCount; j < len(resultInt); j++ {
			if countTemp != resultInt[j] {
				resultInt = append(resultInt, countTemp) //把第一个元素追加进去
				resultCount++
			}
		}
	}
	return len(resultInt), resultInt
}

// 最长公共前缀
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longesCommonPrefix(strs []string) string {
	bytes := []byte(strs[0])
	var result []byte
outerLoop:
	for i := 0; i < len(bytes); i++ {
		byteCh := bytes[i]
		flat := true
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				break outerLoop
			}
			if strs[j][i] != byteCh { //从字符串strs中获取第j个的第i个字符返回码点
				flat = false
				break outerLoop
			}
		}
		if flat {
			result = append(result, byteCh)
		}
	}
	return string(result)
}

// 获取回文数
func hwString(str string) bool {
	mapBack := make(map[int32]int32)
	mapBack[40] = 41   //()
	mapBack[123] = 125 //{}
	mapBack[91] = 93   //[]
	flat := true
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		inputChar := chars[i]
		value, exist := mapBack[inputChar]

		if exist {
			flat = false
			for j := i; j < len(chars); j++ {
				if value == chars[j] {
					flat = true
				}
			}
		} else {
			continue
		}

	}
	return flat
}

// 获取重复值
func arrGetOne(arr []int) []int {

	//	inputArr := []int{1, 2, 3, 7, -2, 3, 9, 2}
	//	var resultArr []int = arrGetOne(inputArr)
	//	fmt.Println(resultArr)
	mapSave := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		_, exist := mapSave[num]
		if exist {
			mapSave[num] += 1
		} else {
			mapSave[num] = 1
		}
	}
	var resultArr []int
	for key, value := range mapSave {
		if value == 1 {
			resultArr = append(resultArr, key)
		}
	}
	return resultArr
}
