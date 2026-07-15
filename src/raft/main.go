package main

import (
	"raft/node"

	"log"
	"time"
)

func main() {
	log.Printf("hello node")

	HBch := make(chan string)

	node1 := node.Node{HBch}
	
	go node1.Timeout()

	time.Sleep(2 * time.Second)

	HBch <- "signal"
}