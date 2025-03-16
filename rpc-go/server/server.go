package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
)

// Args holds the arguments passed to the RPC method
type Args struct {
    A, B int
}

// Arith provides an RPC service with arithmetic operations
type Arith int

// Multiply is an RPC method that multiplies two integers
func (t *Arith) Multiply(args *Args, reply *int) error {
    if args.A == 0 || args.B == 0 {
        return errors.New("arguments must be non-zero")
    }
    *reply = args.A * args.B
    return nil
}

// Add is an RPC method that adds two integers
func (t *Arith) Add(args *Args, reply *int) error {
    if args.A == 0 || args.B == 0 {
        return errors.New("arguments must be non-zero")
    }
    *reply = args.A + args.B
    return nil
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    l, err := net.Listen("tcp", ":1234")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer l.Close()
    fmt.Println("Server listening on port 1234")
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}