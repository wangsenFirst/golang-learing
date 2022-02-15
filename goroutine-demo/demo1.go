package main

import (
	"fmt"
	"strings"
	"sync"
)

/**
一个发送多个接受：发送完直接关闭通道
多个发送一个接受：两个sync.WaitGroup  使用sync.WaitGroup确保都发送完了在关闭通道，零个sync.WaitGroup来控制消费的通道的service
多个发送多个接受：创建两个通道 循环开启到最后一个goroutine 等待几秒再关闭另一个通道  select 另一个通道通道关闭的时候会case一次，来确定通道关闭
*/
// sender: 1, receivers: M
const (
	NumReceivers1 = 20
)

var wg sync.WaitGroup
var strMessages = []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii", "jjj", "kkk"}

func sender(msg []string, ch chan<- string) {
	for i, v := range msg {
		ch <- v
		if i == len(msg)-1 {
			close(ch)
		}
	}

}

func receiver(ch chan string) {
	defer wg.Done()

	for value := range ch {
		fmt.Println(strings.ToUpper(value)) // process messages
	}
}

func main() {
	//rand.Seed(time.Now().Unix())

	ch := make(chan string, 20)

	go sender(strMessages, ch)

	wg.Add(NumReceivers1)
	// 开启20个接受
	for i := 0; i < NumReceivers1; i++ {
		go receiver(ch)
	}
	wg.Wait()
}
