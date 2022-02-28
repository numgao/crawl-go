package main

import "fmt"
//import "io"

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error){
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main (){
	var c ByteCounter
	println(c)
	var b, _ = c.Write([]byte("hello"))
	println("b=", b)
	println(c)
	println("***************")
	c =0
	var name = "Gao Yu"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

