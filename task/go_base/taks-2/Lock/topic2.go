package Lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sw sync.WaitGroup
	var count int64 = 0 // 修改为 int64 类型
	for i := 0; i < 10; i++ {
		sw.Add(1) //相当于一个计数器
		go func() {
			defer sw.Done() // 会对add减去一 直到等于0 	sw.Wait() 结束等待响应
			add1000Ao(&count)
		}()
	}
	sw.Wait()
	fmt.Println("最终结果：", count)
}

func add1000Ao(count *int64) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(count, 1)
	}
}
