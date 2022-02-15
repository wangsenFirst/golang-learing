package main

import (
	"context"
	"fmt"
	"sync"
)

//const maxGoroutineNum = 10
/**
多个发送一个接受
*/

var wgSender sync.WaitGroup
var wgReceiver sync.WaitGroup

type ProcessMessage func(ctx context.Context, originMsg *Person) (resMsg *Person)

type Person struct {
	Name string
	Job  string
}

func main() {
	ctx := context.Background()
	firstPerson := &Person{Name: "John", Job: "worker"}
	secondPerson := &Person{Name: "Tom", Job: "engineer"}
	originalMsg := []*Person{firstPerson, secondPerson}
	ch := make(chan *Person, len(originalMsg))
	// 开启两个goroutine 往通道里面放值
	StartSend(ctx, originalMsg, ch, process)

	wgReceiver.Add(1)
	var resList []*Person
	// 开启一个goroutine 取值
	go func() {
		resList = StartReceive(ch)
	}()
	wgSender.Wait()
	// 发送结束关闭通道  select case中退出for循环
	close(ch)
	wgReceiver.Wait()
	fmt.Println(resList[0].Job, resList[1].Job)
}

func StartSend(ctx context.Context, originalMsgs []*Person, ch chan *Person, process ProcessMessage) {
	for _, originalMsg := range originalMsgs {
		wgSender.Add(1)
		go Sender(ctx, originalMsg, ch, process)
	}
}

func StartReceive(ch chan *Person) (resList []*Person) {
	defer wgReceiver.Done()
	resList = Receiver(ch)
	return
}

func Sender(ctx context.Context, originalMsg *Person, ch chan *Person, process ProcessMessage) {
	defer wgSender.Done()
	ch <- process(ctx, originalMsg)

}

func Receiver(ch chan *Person) (eleList []*Person) {
	for {
		select {
		case ele, isOpened := <-ch:
			if !isOpened {
				return eleList
			}
			eleList = append(eleList, ele)
		default:
		}

	}
}

func process(ctx context.Context, originMsg *Person) (resMsg *Person) {
	originMsg.Job = "good " + originMsg.Job
	return originMsg

}
