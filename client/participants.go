// s2
package main

/*import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	proto "chit-chat/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientName string
var grpcClient proto.ChatServiceClient

func connection() *grpc.ClientConn {
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

/*func NewParticipantJoined(name string) {
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
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Stream closed: %v", err)
			return
		}

		//print message
		fmt.Printf("\n[%s]: %s>n ", msg.Sender, msg.Content)
		// log message
		fmt.Printf("Reseived from %s: %s", msg.Sender, msg.Content)
	}
}

func publishMsg() {
	// max length 128
	// UTF-8 encoded string
	if len(content) > 128 {
		fmt.Println("Error: The message must be under 128 characters")
		return
	}

	// check if emty
	if content == "" {
		return
	}

	log.Printf("Send message: %s", content)

	// make message
	msg := &proto.Message{
		Sender:  clientName,
		Content: content,
	}

	// Send to server
	_, err := grpcClient.SendMessage(context.Background(), msg)
	if err != nil {
		log.Printf("The Message failed to sand: %v", err)
		fmt.Println("Error: The massage could not send")
	}
}

func leaveChat(conn *grpc.ClientConn) {
	log.Println("Leaving the chat...")
	conn.Close()
	log.Println("Goodbye!")
}

func readUserInput() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n=Chit Chat=")
	fmt.Println("Write your message")
	fmt.Println("Write 'exit' to leave")
	fmt.Println("========\n")

	for {
		fmt.Print(">")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "exit" {
			break
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

	NewParticipantJoined()

	readUserInput()
}
*/
