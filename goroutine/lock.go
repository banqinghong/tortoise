package main

import (
	"fmt"
	"sync"
)

/*
下面的程序在没有加锁的情况下，最后输出的结果并不会是1000000，
出现的原因就是count++这个动作不是原子动作，
多个goroutine在操作count的时候会相互干扰，导致每次输出的结果并不一致
*/
func main() {
	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("count =", count)
}

// 不会顺序输出，因为goroutine创建和调度不是同时发生的,且调度具有一定的随机性
func sort() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
