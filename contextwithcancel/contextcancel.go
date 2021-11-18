package main

import (
	"context"
	"fmt"
	"net"
	"syscall"
	"time"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
    sync := make(chan struct{})

     go func() {
        defer func() { sync <- struct{}{} }()

        var d net.Dialer
        d.Control = func(_, _ string, _ syscall.RawConn) error {
            time.Sleep(time.Second)
            return nil
        }
        conn, err := d.DialContext(ctx, "tcp", "10.0.0.1:80")
        if err != nil {
            fmt.Println(err)
            return
        }

        conn.Close()
       fmt.Println("connection did not time out")
    }()

     cancel()
    <-sync

    if ctx.Err() != context.Canceled {
        fmt.Printf("expected canceled context; actual: %q", ctx.Err())
    }
}