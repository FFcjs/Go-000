package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main(){
	conn,err:=net.Dial("tcp","localhost:9999")
	if err!=nil{
		log.Println("connect failed")
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(conn)
	rd:=bufio.NewReader(conn)
	go func(){
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				log.Printf("read error: %v\n", err)
				break
			}
			log.Printf("receive msg:%s",line)
		}
	}()
	for{
		line,err := reader.ReadString('\n')
		if err!=nil{
			log.Printf("read from keyboard fail,err:%v",err)
			break
		}
		_,err=wr.Write([]byte(line))
		if err!=nil{
			log.Printf("write fail,err:%v",err)
			break
		}
		err=wr.Flush()
		if err!=nil{
			log.Printf("flush fail,err:%v",err)
			break
		}
		log.Printf("send msg: %s",line)
	}
}
