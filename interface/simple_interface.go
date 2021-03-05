package main

import "fmt"

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	Num int
}

func (s *Simple) Get() int {
	return s.Num
}

func (s *Simple) Set(i int) {
	s.Num = i
}

func OpSimple(s Simpler) {
	fmt.Println(s.Get())
	s.Set(333)
	fmt.Println(s.Get())
}

func main() {
	simple := Simple{1}
	var simpler Simpler
	simpler = &simple
	OpSimple(simpler)
}
