package main
import (
	"os"
	"net"
	"fmt")

func main (){
	name := os.Args[1]
	addr := net.ParseIP(name)
	fmt.Println("the address is ", addr.String())

}
