package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
   conn,err :=	net.DialTimeout("tcp","10.0.01:http",time.Second*5)
   if err == nil {
	conn.Close()
	fmt.Println("connection did not time out")
}
nErr, ok := err.(net.Error)
if !ok {
	fmt.Println(err)
}
if !nErr.Timeout() {
	fmt.Println("error is not a timeout")
}
}