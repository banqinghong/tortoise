package main

import (
	"fmt"
	"sync"
)

/*
下面程序输出结果并不是有序的
原因是goroutine在创建完成之后并不会马上执行，而是会等待调度。
且goroutine调度也是随机的
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println("i = ", n)
		}(i)
	}
	wg.Wait()
}
