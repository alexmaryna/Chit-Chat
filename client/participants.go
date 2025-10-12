// s2
package main

/*import (
	"context"
	"fmt"
	"log"
	"bufio"
	"os"

	pb "chit-chat/grpc"
	proto "chit-chat/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientName string
var grpcClient proto.NewChatServiceClient

func connection() *grpc.ClientConnection {
	log.Println("Connecting to server...")

	conn, err := grpc.NewClient(
		"localhost:5050",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	log.Println("Connected")
	return conn
}

func NewParticipantJoined(name string) {
	log.Printf("%s is trying to join the chat...", clientName)

	// start stream
	stream, err := grpcClient.ReceiveMessages(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Could not join the chat: %v", err)
	}

	log.Println("Joined the chat succesfully")

	go receiveMessage(stream)
}

func receiveMessage(stream proto.ChatService_ReceiveMessagesClient) {
	for {

	}
}

type Participant struct {
	name string
}

func main() {
	// get name

	// connect
	conn, err := grpc.NewClient("localhost:5050",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)

	//subscribe
	stream, err := client.ReceiveMessages(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Could not subscribe: %v", err)
	}

	fmt.Println("Connnected...")

	//receive
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Printf("Could closed: %v", err)
			return
		}
		fmt.Printf("[%d] %s: %s\n", message.GetLogicTime(), message.GetSender(), message.GetContent())
	}
	//timestamps in logs

	//connection to the server

	//
}

func leave() {}
func publishMsg() {
	// max length 128
	// UTF-8 encoded string
}

*/