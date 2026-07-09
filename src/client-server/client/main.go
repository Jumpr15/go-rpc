package main

import (
	"log"
	"net/rpc"
)

type Args struct {
}

func main() {
	var reply string
	args := Args{}

	client, err := rpc.DialHTTP("tcp", "localhost:2233")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}
	err = client.Call("ServerRes.GiveServerRes", args, &reply)
	if err != nil {
		log.Fatal("RPC Call error: ", err)
	}
	log.Printf(reply)
}