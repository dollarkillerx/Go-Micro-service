/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-3
* Time: 上午10:17
* */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "name", "hello")
	go watch(ctx,"[监控1]")
	go watch(ctx,"[监控2]")
	go watch(ctx,"[监控3]")

	time.Sleep(time.Second * 10)
	fmt.Println("通知停止监控")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context,name string)  {
	for {
		select {
		case <- ctx.Done():
			fmt.Println(name,"监控退出,停止了....")
			return
		default:
			fmt.Println(ctx.Value("name"))
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}