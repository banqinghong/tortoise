package main

import (
	"errors"
	"fmt"
)

type Driver interface {
	Driving() error
}

type Person string

func (p Person) Driving() (err error) {
	//fmt.Println("人是可以开车的")
	return nil
}

type Dog string

func (d Dog) Driving() (err error) {
	//fmt.Println("狗不会开车")
	return errors.New("111")
}

func IsDriver(d Driver) bool {
	err := d.Driving()
	if err != nil {
		return false
	}
	return true
}

func main() {
	p1 := Person("banqinghong")
	qiuqiu := Dog("dog")
	if IsDriver(p1) {
		fmt.Printf("%s 会开车\n", p1)
	}
	if IsDriver(qiuqiu) {
		fmt.Printf("%s 不会开车\n", qiuqiu)
	}
}
