package node

import (
	"time"
	"math/rand"
	"log"
)

type Node struct {
	isLeader bool
	HeartbeatCh chan string
}

func (n Node) Timeout() {
	for i := 0; i < 99; i++ { // hardcoded no. loops
		rand.Seed(time.Now().UnixNano())
		s := rand.Intn(5) // first generate a randomized heartbeat timeout duration

		select {
		case sig := <- n.HeartbeatCh:
			log.Printf("Received heartbeat", sig)
		case <- time.After(time.Duration(s) * time.Second):
			log.Fatal("No heartbeat received. Initializing leader re-election")
		}
	}
}

func SendElectionSig