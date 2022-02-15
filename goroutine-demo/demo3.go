package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type ProcessMsg func(ctx context.Context, originMsg interface{}) (resMsg interface{})

var (
	wg3 sync.WaitGroup
)

func senders3(i int, msg []string, stopCh chan struct{}, ch chan string) {
	for _, v := range msg {
		ch <- v
	}
	if i == -1 {
		//循环到最后一次等待其他 goutine执行完再关闭通道stopCh
		time.Sleep(time.Millisecond)
		//close(stopCh)
	}

}

func receivers3(ctx context.Context, stopCh chan struct{}, ch chan string, process ProcessMsg) {
	defer wg3.Done()
	time.Sleep(time.Second * 3)
	for {
		select {
		// stopCh通道关闭后胡进入这个case
		case <-stopCh:
			// stopCh通道关闭后胡进入这个case 然后退出for循环, 为什么不直接关闭ch通道？ 因为不知道ch通道什么时候不需要在放值了
			println("stopCh")
			return

		case value := <-ch:
			resMsg := process(ctx, value)
			fmt.Println(resMsg)
		default:
		}
	}
}

func processMsg(ctx context.Context, originMsg interface{}) (resMsg interface{}) {
	msg := originMsg.(string)
	msg = strings.ToUpper(msg)
	return msg
}

func main() {
	ctx := context.Background()

	ch := make(chan string, 100)
	stopCh := make(chan struct{})

	msg := [][]string{{"abc", "def", "ghi"}, {"123", "234", "345"}, {"a1a", "b2b", "c3c"}}

	length := len(msg)
	// 开启三个 goroutine 去往通道里面放值
	for i := 0; i < length; i++ {
		// 循环三次
		j := i
		if i == length-1 {
			// 循环到最后一次
			j = -1
		}
		// i 0 1 2
		// j 0 1 -1
		go senders3(j, msg[i], stopCh, ch)
	}
	// 添加三个等待
	wg3.Add(length)
	// 开启多个 goroutine去通道取值
	for i := 0; i < length; i++ {
		go receivers3(ctx, stopCh, ch, processMsg)
	}
	wg3.Wait()

}
