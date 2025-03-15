package main

import (
	"fmt"
	"net/rpc"
)

// Args holds the arguments passed to the RPC method
type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer client.Close()

	args := Args{A: 2, B: 3}
	var reply int
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		fmt.Println("Error calling Arith.Multiply:", err)
		return
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Test with invalid arguments
	args = Args{A: 0, B: 3}
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		fmt.Println("Error calling Arith.Multiply with invalid args:", err)
		return
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
