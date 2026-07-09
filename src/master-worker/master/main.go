package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	Req_string string
}

func main() {
	var reply string
	Req_string := "request 1"
	args := Args{Req_string}

	client, err := rpc.DialHTTP("tcp", "localhost:2233")
	if err != nil {
		log.Fatal("Dialing error: ", err)
	}

	err = client.Call("WorkerRes.GiveWorkerRes", args, &reply)
	if err != nil {
		log.Fatal("RPC Call error: ", err)
	}
	log.Printf(reply)
}