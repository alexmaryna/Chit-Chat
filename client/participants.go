// s2
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	proto "chit-chat/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientName string
var grpcClient proto.ChatServiceClient

func connection() *grpc.ClientConn {
	log.Println("Connecting to server...")

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	return conn
}

func NewParticipantJoined(name string) {
	// start stream
	stream, err := grpcClient.ReceiveMessages(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("Could not start receiving meassages: %v", err)
	}

	go receiveMessage(stream)

	if _, err := grpcClient.Join(context.Background(), &proto.User{Name: name}); err != nil {
		log.Fatalf("Could not join the chat: %v", err)
	}

	//log.Println("Joined the chat succesfully")
}

func receiveMessage(stream proto.ChatService_ReceiveMessagesClient) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Connection closed: %v", err)
			return
		}

		// print message
		fmt.Printf("\n[%d] %s: %s\n> ", msg.GetLogicalTime(), msg.GetSender(), msg.GetContent())

		// log message
		//log.Printf("Received Time: %d, From %s Message: %s", msg.GetLogicalTime(), msg.GetSender(), msg.GetContent())
	}
}

func publishMsg(content string) {
	content = strings.TrimSpace(content)

	// check if emty
	if content == "" {
		return
	}

	if len(content) > 128 {
		fmt.Println("Error: The message must be under 128 characters")
		return
	}

	//log.Printf("Send message: %s", content)

	// make message
	msg := &proto.Message{
		Sender:  clientName,
		Content: content,
	}

	// Send to server
	if _, err := grpcClient.SendMessage(context.Background(), msg); err != nil {
		//log.Printf("The Message failed to sand: %v", err)
		fmt.Println("Error: The massage could not send")
	}
}

func leaveChat(conn *grpc.ClientConn) {
	//log.Println("Leaving the chat...")
	_, _ = grpcClient.Leave(context.Background(), &proto.User{Name: clientName})
	_ = conn.Close()
}

func readUserInput() {
	scanner := bufio.NewScanner(os.Stdin)

	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n== Chit Chat ==")
	fmt.Println("Write your message and press Enter")
	fmt.Println("Write 'exit' to leave")
	fmt.Println("==================================")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		input := scanner.Text()

		if strings.EqualFold(input, "exit") {
			return
		}

		publishMsg(input)
	}
}

func main() {
	log.SetFlags(log.Ltime)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <your_name>")
		os.Exit(1)
	}

	clientName = os.Args[1]

	conn := connection()
	defer leaveChat(conn)

	grpcClient = proto.NewChatServiceClient(conn)

	NewParticipantJoined(clientName)

	readUserInput()
}
