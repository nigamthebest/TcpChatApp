package main

import (
	"testing"
	"net"
	"fmt"
	"bufio"
)

func init() {
	startHub()
}
func TestHandelMessage(t *testing.T) {

	conn, error := net.Dial("tcp", "127.0.0.1:8080")
	if error != nil {
		panic(error)
	}
	text := "{'messageType':'ID','senderId' :123,'receiverIds':[1,2],'messageBody':'string'}"
	fmt.Fprintf(conn, text)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: "+message)

}