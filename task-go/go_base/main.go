package go_base

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flo", "floght"}
	fmt.Println(longesCommonPrefix(strs))

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
			if !longesCommonPrefixChar(strs[j], i, byteCh) {
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

func longesCommonPrefixChar(targe string, targeIndex int, ch byte) bool {
	bytes := []byte(targe)
	if bytes[targeIndex] == ch {
		return true
	}
	return false
}
