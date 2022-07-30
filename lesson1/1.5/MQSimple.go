package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	baseCtx := context.Background()
	// 设置main的超时
	timeoutCtx, cancelFunc := context.WithTimeout(baseCtx, 10*time.Second)
	defer cancelFunc()

	producerWith1Second(ch, timeoutCtx)
	consumerWith1Second(ch, timeoutCtx)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("MAIN process exit")
	}
}

// producerWith1Second 每1s生产一个数据
func producerWith1Second(ch chan<- int, ctx context.Context) {
	go func(ch chan<- int, ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("PRODUCER child process interrupt")
				return
			default:
				// 生成随机数来发送
				rand.Seed(time.Now().UnixNano())
				sendValue := rand.Intn(100)
				ch <- sendValue
				fmt.Printf("produce val: %d\n", sendValue)
			}
		}
	}(ch, ctx)
}

// consumerWith1Second 每1s消费一个数据
func consumerWith1Second(ch <-chan int, ctx context.Context) {
	go func(ch <-chan int, ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("CONSUMER child process interrupt")
				return
			default:
				fmt.Printf("consumer value: %d\n", <-ch)
			}
		}
	}(ch, ctx)
}
