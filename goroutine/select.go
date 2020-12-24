package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SayHello(ch chan bool) {
	time.Sleep(300 * time.Millisecond)
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(999)
	fmt.Println("[hello] num = ", randNum)
	if randNum%2 == 0 {
		ch <- true
	}
	close(ch)
}

func SayHi(ch chan bool) {
	time.Sleep(200 * time.Millisecond)
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
	time.Sleep(150 * time.Millisecond)
	select {
	case result1 := <-chHello:
		fmt.Println("hello...")
		if result1 {
			fmt.Println("hello world")
		}
	case result2 := <-chHi:
		fmt.Println("hi...")
		if result2 {
			fmt.Println("hi world")
		}
		//default:
		//	fmt.Println("nothing...")
	}
	time.Sleep(1 * time.Second)
}
