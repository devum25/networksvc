package main

import (
	"context"
	"fmt"
	"net"
	"syscall"
	"time"
)

func main(){

	dl := time.Now().Add(time.Second*5)
	ctx,cancel := context.WithDeadline(context.Background(),dl)

	defer cancel()

	var d net.Dialer //Dialcontext is a method on dialer

	d.Control = func(_, _ string, _ syscall.RawConn) error {
        // Sleep long enough to reach the context's deadline.
        fmt.Println("delaying connection")
        time.Sleep(5*time.Second + time.Millisecond)
        return nil
    }
    conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
    if err == nil {
        conn.Close()
       fmt.Println("connection did not time out")
    }
    nErr, ok := err.(net.Error)
    if !ok {
        fmt.Println(err)
    } else {
        if !nErr.Timeout() {
           fmt.Printf("error is not a timeout: %v", err)
        }
    }
     if ctx.Err() != context.DeadlineExceeded {
        fmt.Printf("expected deadline exceeded; actual: %v", ctx.Err())
    }
    if ctx.Err() == context.DeadlineExceeded {
        fmt.Printf("expected deadline exceeded; actual: %v", ctx.Err())
    }
}

