package main

import (
	"net"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"bufio"
	"encoding/json"
)

type Client struct {
	user_id int
	reader  *bufio.Reader
	writer  *bufio.Writer
}

type Message struct {
	MessageType string `json:"messageType"`
	SenderId    int `json:"senderId"`
	ReceiverIds []int `json:"receiverIds"`
	MessageBody string `json:"messageBody"`
}

func main() {
	startHub()
}

func startHub() {
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


func handleClient(connection net.Conn, clientMap map[int]Client) {
	fmt.Println("In  handleClient")
	user_id := rand.Int()
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)
	client := Client{user_id, reader, writer}
	var clientsMap map[int]Client = make(map[int]Client)
	clientsMap[user_id] = client
	for {
		// Read message JSON
		messageBytes, _ := reader.ReadBytes('\n');
		//convert to JSON to obj
		message := Message{}
		error := json.Unmarshal(messageBytes, &message)
		if (error != nil) {
			fmt.Println("error in Unmarshal", error)
		}

		fmt.Println("message ", message)
		switch  message.MessageType{
		case "ID":
			handelId(client)
		case "LIST":
			handleList(client, clientsMap)
		case "RELAY":
			handleRelay(message, client, clientsMap)
		default:
			fmt.Println("Unknown mesaage try 'ID' 'LIST'")
		}
	}
}

func handleRelay(message Message, client Client, clientsMap map[int]Client) {
	validateSenderId(message.SenderId, clientsMap)
}
func validateSenderId(user_id int, clientsMap map[int]Client) {

}
func (client *Client) Write(data string) {
	client.writer.WriteString(data)
	client.writer.Flush()
}


func handelId(client Client) {
	fmt.Println("Handeling ID message")
	client.Write("your Id is " + strconv.Itoa(client.user_id) + "\n")
}

func handleList(client Client, clientsMap map[int]Client) {
	fmt.Println("Handeling LIST message")
	userIds := make([]string, 0, len(clientsMap))
	for k := range clientsMap {
		userIds = append(userIds, strconv.Itoa(k))
	}
	// send new string back to client
	client.Write("connectedIds are " + strings.Join(userIds, ",") + "\n")
}
