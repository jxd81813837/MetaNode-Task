package Channel

import (
	"fmt"
	"sync"
)

func main_channel_2() {
	chanCh := make(chan int, 10)
	var sw sync.WaitGroup
	sw.Add(1)
	go func() {
		defer sw.Done()
		for i := 1; i <= 100; i++ {

			chanCh <- i
			fmt.Println("发送数据：", i)
		}
		close(chanCh)
	}()
	sw.Add(1)
	go func() {
		defer sw.Done()
		for i := range chanCh {
			fmt.Println("接收数据：", i)
		}
	}()

	sw.Wait()
}
