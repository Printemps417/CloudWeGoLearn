package main

import (
	"github.com/cloudwego/kitex/server"
	studentservice "kitex.demo/kitex_gen/demo/studentservice"
	"log"
	"net"
)

func main() {
	//svr := studentservice.NewServer(new(StudentServiceImpl))
	//addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	//svr := studentservice.NewServer(studentservice, server.WithServiceAddr(addr))
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	svr := studentservice.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
