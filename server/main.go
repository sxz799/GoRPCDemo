package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"server/server"
)

func main() {
	err := rpc.RegisterName("MathService", new(server.MathService))
	if err != nil {
		log.Println("rpc register fail ", err.Error())
		return
	}
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, err2 := l.Accept()
		if err2 != nil {
			log.Println("jsonrpc.Serve error: accept:", err2.Error())
			return
		}
		log.Println(conn.RemoteAddr())
		//json rpc
		go jsonrpc.ServeConn(conn)
	}
}
