package main

import "fmt"

type BaseStruct struct{}

type Component interface {
	ChildDo()
}

// ChildDo 执行子组件
func (bs *BaseStruct) ChildDo() {
	fmt.Println("child")
}

type TestStruct struct {
	BaseStruct
}

func (ts *TestStruct) Run() {
	ts.ChildDo()
}

func main() {
	testStruct := &TestStruct{}
	testStruct.Run()
}
