package Channel

import (
	"fmt"
	"sync"
)

func main_cha1() {
	//建立一个发送整数的通道
	sendCh := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Println("发送数据：", i)
			sendCh <- i
		}
		//发送完成
		close(sendCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range sendCh {
			fmt.Println("接收数据：", i)
		}
	}()
	wg.Wait()
}
