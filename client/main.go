package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := "{\"a\":5,\"b\":2}"
	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error:", err)
	}
	fmt.Printf("MathService.Add: %s   %d\n", args, reply)

	err = client.Call("MathService.Sub", args, &reply)
	if err != nil {
		log.Fatal("MathService.Sub error:", err)
	}
	fmt.Printf("MathService.Sub: %s   %d\n", args, reply)
}
