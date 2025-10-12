package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "chit-chat/grpc"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients []pb.ChatService_ReceiveMessagesServer
}

func (s *ChatServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("%s: %s", msg.Sender, msg.Content)

	for _, stream := range s.clients {
		stream.Send(msg)
	}

	return &pb.Ack{Info: "Message received by server"}, nil
}

func (s *ChatServer) ReceiveMessages(empty *pb.Empty, stream pb.ChatService_ReceiveMessagesServer) error {
	s.mu.Lock()
	s.clients = append(s.clients, stream)
	s.mu.Unlock()

	<-stream.Context().Done()
	return nil
}

// S5–S6 til join/leave beskeder
func broadcast() {
	// if they leave or join send message
	// Send publish message to active users
	// must include message and timestamp
}




func (s *ChatServer) Leave(ctx context.Context, user *pb.user) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()


	logicalTime := time.Now().UnixMilli()

	
	leaveMsg := &pb.Message{
		Sender:      "Server",
		Content:     fmt.Sprintf("Participant %s left Chit Chat at logical time %d", user.Name, logicalTime),
		logicalTime: logicalTime,
	}

	for _, stream := range s.clients {
		stream.Send(leaveMsg)
	}

	log.Printf("Participant %s left Chit Chat at %d", user.Name, logicalTime)
	return &pb.Ack{Info: "Leave message broadcasted"}, nil
}




func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &ChatServer{})

	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
