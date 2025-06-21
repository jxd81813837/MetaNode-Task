package main

import "MetaNode-Task/study"

func main() {
	study.QueryBlock()

	//bytesSlice := [32]byte{100, 101, 109, 111, 95, 74, 88, 68, 95, 118, 97, 108, 117, 101, 49, 49, 49, 49, 49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//bytesSlice1 := [32]byte{100, 101, 109, 111, 95, 74, 88, 68, 95, 118, 97, 108, 117, 101, 49, 49, 49, 49, 49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//hexStr := "64656d6f5f4a58445f6b65790000000000000000000000000000000000000000"
	// 解码 hex 字符串为字节切片
	//data, err := hex.DecodeString(hexStr)
	//var key [32]byte
	//copy(key[:], data)
	//fmt.Println("Key:", string(key[:]))
	//fmt.Println(string(bytesSlice[:]))
	//fmt.Println(string(bytesSlice1[:]))
}
