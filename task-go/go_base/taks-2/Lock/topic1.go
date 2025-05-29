package Lock

import (
	"fmt"
	"sync"
)

// 全局变量
var lock sync.Mutex

func main_lock_1() {
	var sw sync.WaitGroup
	count := 0
	for i := 0; i < 10; i++ {
		sw.Add(1) //相当于一个计数器
		go func() {
			defer sw.Done() // 会对add减去一 直到等于0 	sw.Wait() 结束等待响应
			add1000(&count)
		}()
	}
	sw.Wait()
	fmt.Println("最终结果：", count)
}

func add1000(count *int) {
	lock.Lock()
	for i := 0; i < 1000; i++ {
		*count++
	}
	lock.Unlock()
}
