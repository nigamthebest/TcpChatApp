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
	user_id   int
	reader    *bufio.Reader
	writer    *bufio.Writer
	ClientMap *map[int]Client
}

type Message struct {
	MessageType string `json:"messageType"`
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
	client := Client{user_id, reader, writer, &clientMap}
	clientMap[user_id] = client
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
			handleList(client)
		case "RELAY":
			handleRelay(message, client)
		default:
			fmt.Println("Unknown mesaage try 'ID' 'LIST'")
		}
	}
}

func handleRelay(message Message, client Client) {
	for _, receiver := range message.ReceiverIds {
		receiverMap := *client.ClientMap
		receivingClient := receiverMap[receiver]
		receivingClient.Write(message.MessageBody)
		client.Write("Your Message " + message.MessageBody +" Sent to "+ strconv.Itoa(receivingClient.user_id) + "\n")
	}

}



func (client *Client) Write(data string) {
	client.writer.WriteString(data)
	client.writer.Flush()
}


func handelId(client Client) {
	fmt.Println("Handeling ID message")
	client.Write("your Id is " + strconv.Itoa(client.user_id) + "\n")
}

func handleList(client Client) {
	fmt.Println("Handeling LIST message", len(*client.ClientMap))
	userIds := make([]string, 0, len(*client.ClientMap))
	for k := range *client.ClientMap {
		if (k != client.user_id ) {
			userIds = append(userIds, strconv.Itoa(k))
		}
	}
	// send new string back to client
	client.Write("connectedIds are " + strings.Join(userIds, ",") + "\n")
}
