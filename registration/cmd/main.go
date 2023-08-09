package main

import (
	"net/rpc"
	"github.com/gin-gonic/gin"
)

func main() {

	type RpcServer struct{}

	func (r *RpcServer) getPatientInfo(resp *int) error {
		*resp = 14
		return nil
	}

	rpc.Register(new(RpcServer))
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return err
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err)
		continue
	}
	go rpc.ServeConn(conn)
}
