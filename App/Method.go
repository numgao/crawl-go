package main

import "fmt"

type Method_test struct{
	name string
	age int
}

func (b Method_test) test() {
	fmt.Println("hte test result")
}

func main(){
	a := Method_test{
		"gao",
		20,
	}
	fmt.Println(a.name, a.age)
	a.test()
}
