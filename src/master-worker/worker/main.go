package main

import (
	"fmt"
	"time"
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	Req_string string
}

type WorkerRes string

func (w *WorkerRes) GiveWorkerRes(args *Args, reply *string) error {
	time.Sleep(time.Second * 2)

	*reply = "Worker response"
	return nil
}

func main() {
	port_flag := flag.Int("port", 1234, "The port which the worker server listens on")
	flag.Parse()

	worker_server := new(WorkerRes)

	rpc.Register(worker_server)
	rpc.HandleHTTP()

	port := fmt.Sprintf(":%d", *port_flag)

	l, e := net.Listen("tcp", port)
	if e != nil {
		log.Fatal("Listening error: ", e)
	}

	http.Serve(l, nil)
}