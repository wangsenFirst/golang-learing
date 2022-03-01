package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("ok")
}
func main() {
	go func() {
		//1在这里需要你写算法
		//2要求每秒钟调用一次proc函数
		//3要求程序不能退出
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					fmt.Println(time.Now())
					proc()
				}()
			}
		}

	}()

	select {}
}
