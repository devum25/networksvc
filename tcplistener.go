package main

import (
	"fmt"
	"net"
)

func main(){
  listener, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        fmt.Println(err)
    }
     defer func() { _ = listener.Close() }()

	 fmt.Printf("bound to %v\n",listener.Addr())

	 for{

		conn,err := listener.Accept()
		if err != nil{
			 fmt.Println(err)
			 return
		}

		go func(c net.Conn){
			fmt.Printf("recieved request from %v\n",c.RemoteAddr())
			defer c.Close()
		}(conn)
	 }

}