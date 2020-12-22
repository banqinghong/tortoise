package main

import (
	"fmt"
)

// 交替打印字母和数字
//func main()  {
//	numChan := make(chan bool)
//	letChan := make(chan bool)
//	exitChan := make(chan bool)
//	endNum := 5
//	startLetter := 'A'
//	go func () {
//		for i := 1; i <= endNum; i++ {
//			<- numChan
//			fmt.Printf("%d", i)
//			letChan <- true
//		}
//	}()
//	go func() {
//		for i := 1; i <= endNum; i++ {
//			<- letChan
//			fmt.Printf("%c\n", startLetter)
//			startLetter++
//			if i < endNum {
//				numChan <- true
//			}
//		}
//		exitChan <- true
//	}()
//	numChan <- true
//	<- exitChan
//}

func number(numChan, letterChan chan bool, endNum int) {
	for i := 1; i <= endNum; i++ {
		<-numChan
		fmt.Printf("%d", i)
		letterChan <- true
	}
}

func letter(numChan, letterChan, exitChan chan bool, endNum int) {
	startLetter := 'A'
	for i := 1; i <= endNum; i++ {
		<-letterChan
		fmt.Printf("%c\n", startLetter)
		startLetter++
		if i < endNum {
			numChan <- true
		}
	}
	exitChan <- true
}

// 交替打印数字以及字母
func main() {
	numChan := make(chan bool)
	letChan := make(chan bool)
	exitChan := make(chan bool)
	endNum := 9

	go number(numChan, letChan, endNum)
	go letter(numChan, letChan, exitChan, endNum)

	numChan <- true
	<-exitChan
}
