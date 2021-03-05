package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SayHello以及SayHi函数生成随机数并判断奇偶性，奇数返回false，偶数返回true
// 主函数根据goroutine返回结果打印
func SayHello(ch chan bool) {
	//time.Sleep(300 * time.Millisecond)
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(999)
	fmt.Println("[hello] num = ", randNum)
	if randNum%2 == 0 {
		ch <- true
	}
	close(ch)
}

func SayHi(ch chan bool) {
	//time.Sleep(300 * time.Millisecond)
	rand.Seed(time.Now().Unix() - 100)
	randNum := rand.Intn(999)
	fmt.Println("[hi] num = ", randNum)
	if randNum%2 == 0 {
		ch <- true
	}
	close(ch)
}

func main() {
	chHello := make(chan bool)
	chHi := make(chan bool)
	go SayHello(chHello)
	go SayHi(chHi)
	// 不加等待时间，select会因为两个case语句都没有准备好，而直接执行default语句
	// 添加等待时间之后就看两个协程谁先完成了
	time.Sleep(150 * time.Millisecond) // 等待协程先运行完成
	select {
	case result1 := <-chHello:
		//time.Sleep(400 * time.Millisecond)
		fmt.Println("hello...")
		if result1 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	case result2 := <-chHi:
		//time.Sleep(400 * time.Millisecond)
		fmt.Println("hi...")
		if result2 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	default:
		fmt.Println("nothing...")
	}
	time.Sleep(1 * time.Second)
}
