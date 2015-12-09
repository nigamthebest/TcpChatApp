package main

import (
	"net"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"encoding/json"
)

type Client struct {
	user_id int
	Conn    net.Conn
}

type Message struct {
	messageType string
	senderId    int
	receiverIds []int
	messageBody string
}


func main() {
	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")
	// accept connection on port
	clientsMap := make(map[int]Client)
	for {
		conn, _ := ln.Accept()
		go handleClient(conn, clientsMap)
	}
}
func handleClient(con net.Conn, clientMap map[int]Client) {
	for {
		// will listen for message to process ending in newline (\n)
		decoder := json.NewDecoder(con)
		// output message received
		var message Message
		if err := decoder.Decode(&message);
		err != nil {
			panic(err)
		}
		go handelMessage(con, message, clientMap)
	}
}

func handelMessage(con net.Conn, message Message, clientsMap map[int]Client) {
	fmt.Println("In  handelMessage for message", message)
	switch  message.messageType{
	case "ID":
		go handelId(con, clientsMap)
	case "LIST":
		go handleList(con, clientsMap)
	case "RELAY":
		go handleRelay(con, message, clientsMap)
	default:
		fmt.Println("Unknown mesaage try 'ID' 'LIST'")
	}
}

func handleRelay(con net.Conn, message Message, clientsMap map[int]Client) {
}

func handelId(con net.Conn, clientsMap map[int]Client) {
	fmt.Println("Handeling ID message")
	// send new ID back to client
	user_id := rand.Int()
	newClient := Client{user_id, con}
	clientsMap[user_id] = newClient
	con.Write([]byte("your Id is " + strconv.Itoa(user_id) + "\n"))
}

func handleList(con net.Conn, clientsMap map[int]Client) {
	fmt.Println("Handeling LIST message")
	usserIds := make([]string, 0, len(clientsMap))
	for k := range clientsMap {
		usserIds = append(usserIds, strconv.Itoa(k))
	}
	// send new string back to client
	con.Write([]byte("connectedIds are " + strings.Join(usserIds, ",") + "\n"))
}
