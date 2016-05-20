package gencodec

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type MyArith int

func (t *MyArith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func ExampleServerAndClient() {
	rpc.Register(new(MyArith))
	ln, e := net.Listen("tcp", "127.0.0.1:0") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	address := ln.Addr().String()
	defer ln.Close()

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				continue
			}
			go ServeConn(c)

		}
	}()

	client, err := DialTimeout("tcp", address, time.Minute)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	defer client.Close()

	// Synchronous call
	args := &Args{7, 8}
	var reply Reply
	err = client.Call("MyArith.Mul", args, &reply)
	if err != nil {
		fmt.Println("arith error:", err)
	} else {
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply.C)
	}
}
