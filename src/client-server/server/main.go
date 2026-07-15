package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
}

type ServerRes string

func (r *ServerRes) GiveServerRes(args *Args, reply *string) error {

	*reply = "Server response"
	return nil
}

func main() {
	rpcserver := new(ServerRes) // pointer to unallocated serverres type?

	rpc.Register(rpcserver)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", "0.0.0.0:2233")
	if e != nil {
		log.Fatal("Listening error: ", e)
	}

	http.Serve(l, nil)
}