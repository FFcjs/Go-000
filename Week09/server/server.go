package main

import (
	"bufio"
	"context"
	"log"
	"net"
)


func handleConn(conn net.Conn) {
	defer conn.Close()
	MsgChan := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println("receive a conn")
	go handleMessage(ctx, conn, MsgChan)
	defer close(MsgChan)

	rd := bufio.NewReader(conn)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("read error: %v\n", err)
			break
		}
		log.Printf("receive msg:%s",line)
		MsgChan <- line
	}
	log.Printf("handle conn done")
}

func handleMessage(ctx context.Context, conn net.Conn, connChan chan string) {
	wr := bufio.NewWriter(conn)
	for {
		select {
		case <- ctx.Done():
			log.Println("handle message Done")
			return
		case line := <- connChan:
			_,err:=wr.Write([]byte("hello "+line))
			if err!=nil{
				log.Printf("write error: %v\n", err)
				return
			}
			wr.Flush()
			log.Printf("send msg: %s","hello "+line)
		}
	}
}


func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go handleConn(conn)
	}
}