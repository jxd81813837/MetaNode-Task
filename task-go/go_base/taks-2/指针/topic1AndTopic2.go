package 指针

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name string
	Func func() error
}
type Scheduler struct {
	tasks []Task
}

func (s *Scheduler) addScheduler(name string, f func() error) {
	s.tasks = append(s.tasks, Task{
		Name: name,
		Func: f,
	})
}

type TaskResult struct {
	taskName string
	Duration time.Duration
	Error    error
}

func (s *Scheduler) run() []TaskResult {
	var wg sync.WaitGroup
	resultChan := make(chan TaskResult, len(s.tasks))
	taskResult := make([]TaskResult, 0, len(s.tasks))
	for i := range s.tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done() //go func 的代码执行完了 会执行结束 和java 有些区别
			startTime := time.Now()
			task.Func()
			resultChan <- TaskResult{
				taskName: task.Name,
				Duration: time.Since(startTime),
				Error:    nil,
			}
		}(s.tasks[i])
	}
	wg.Wait()
	close(resultChan)
	for t := range resultChan {
		taskResult = append(taskResult, t)
	}
	return taskResult
}

func main_t1() {
	scheduler := Scheduler{}
	scheduler.addScheduler("任务一", func() error {
		fmt.Println("执行任务一....")
		time.Sleep(1 * time.Second)
		return nil
	})
	scheduler.addScheduler("任务二", func() error {
		fmt.Println("执行任务二....")
		time.Sleep(2 * time.Second)
		return nil
	})
	scheduler.addScheduler("任务三", func() error {
		fmt.Println("执行任务三....")
		time.Sleep(3 * time.Second)
		return nil
	})
	//获取结果
	taskResult := scheduler.run()
	for _, result := range taskResult {
		fmt.Printf("任务%s执行完成，耗时%s\n", result.taskName, result.Duration)
	}
}

// 接收一个整数指针
func addTen(count *int) {
	*count += 10
}

// nums []*int 访问每个元素的切片
func mulTow(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2
	}
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
