package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (ss Students) Len() int {
	return len(ss)
}

func (ss Students) Less(i, j int) bool {
	return ss[i].Age < ss[j].Age
}

func (ss Students) Swap(i, j int) {
	tmp := ss[i]
	ss[i] = ss[j]
	ss[j] = tmp
}

func main() {
	students := Students{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 16},
		{Name: "n3", Age: 13},
		{Name: "n4", Age: 13},
	}
	sort.Sort(students)
	fmt.Println(students)
}
