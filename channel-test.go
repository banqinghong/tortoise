package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func Run() {
	i := 0
	for {
		if i == 5 {
			break
		}
		fmt.Println("num: ", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

func exec() {
	go Run()
	time.Sleep(2 * time.Second)
	fmt.Println("all job finished")
}

func main() {
	http.ListenAndServe(":8888", nil)
	startTime := time.Now()

	go exec()

	fmt.Println("start go num: ", runtime.NumGoroutine())
	time.Sleep(10 * time.Second)
	fmt.Println("end go num: ", runtime.NumGoroutine())

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
