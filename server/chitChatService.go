package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "chit-chat/grpc"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	mu       sync.Mutex
	clients  []pb.ChatService_ReceiveMessagesServer
	shutdown chan struct{}
	msgQueue chan *pb.Message // channels use it
}

func (s *ChatServer) messageBroadcastHandle() {
	for msg := range s.msgQueue {
		s.mu.Lock()
		//clientCount := len(s.clients)
		//log.Printf("Broadcatsing message from %s to %d participants", msg.Sender, clientCount)

		for _, stream := range s.clients {
			if err := stream.Send(msg); err != nil {
				log.Printf("Failed to send message to a participant: %v", err)
			}
		}
		s.mu.Unlock()
	}
}

func (s *ChatServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Ack, error) {
	// S3 – Valider beskeden
	if len(msg.Content) > 128 {
		return &pb.Ack{Info: "Message too long (max 128 characters)"}, nil
	}

	// set msg.LogicalTime
	msg.LogicalTime = time.Now().UnixMilli()

	// S4 Broadcast beskeden til alle aktive klienter
	log.Printf("[%d] %s: %s", msg.LogicalTime, msg.Sender, msg.Content)

	// send to channel
	s.msgQueue <- msg

	return &pb.Ack{Info: "Message received"}, nil
}

func (s *ChatServer) ReceiveMessages(empty *pb.Empty, stream pb.ChatService_ReceiveMessagesServer) error {
	s.mu.Lock()
	s.clients = append(s.clients, stream)
	//clientCount := len(s.clients)
	s.mu.Unlock()

	//log.Printf("Participant connected. Total participants: %d", clientCount)

	// Waiting until the participant disconnects
	<-stream.Context().Done()

	// Remove disconnected paticipant from the list
	s.mu.Lock()
	for i, client := range s.clients {
		if client == stream {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			//log.Printf("Participant disconnected. Total participants: %d", len(s.clients))
			break
		}
	}
	s.mu.Unlock()

	return nil
}

// S5 – Når en deltager joiner
func (s *ChatServer) Join(ctx context.Context, user *pb.User) (*pb.Ack, error) {
	logicalTime := time.Now().UnixMilli()
	joinMsg := &pb.Message{
		Sender:      "Server",
		Content:     fmt.Sprintf("Participant %s joined Chit Chat", user.Name),
		LogicalTime: logicalTime,
	}

	log.Printf("Participant %s joined at %d", user.Name, logicalTime)

	// send to channel
	s.msgQueue <- joinMsg

	return &pb.Ack{Info: "Joined"}, nil
}

// S6
func (s *ChatServer) Leave(ctx context.Context, user *pb.User) (*pb.Ack, error) {
	logicalTime := time.Now().UnixMilli()
	leaveMsg := &pb.Message{
		Sender:      "Server",
		Content:     fmt.Sprintf("Participant %s left Chit Chat", user.Name),
		LogicalTime: logicalTime,
	}

	log.Printf("Participant %s left at %d", user.Name, logicalTime)

	s.msgQueue <- leaveMsg

	return &pb.Ack{Info: "Left"}, nil
}

func main() {
	log.SetFlags(log.Ltime) // timestamps tp logs

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ChatServer := &ChatServer{
		shutdown: make(chan struct{}),         // initialiser channel
		msgQueue: make(chan *pb.Message, 100), // buffered channel for messages
	}

	// starts message broadcasting
	go ChatServer.messageBroadcastHandle()

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, ChatServer)

	log.Println("Server started on port 50051")

	// use channel for shutdown signaling
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		close(ChatServer.shutdown)
	}()

	// serve in go routine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Server stopped serving: %v", err)
		}
	}()

	// waiting for shutdown signal via channel
	<-ChatServer.shutdown
	log.Println("Server shutting down...")

	// close msg queue and wait for message broadcast to finish
	close(ChatServer.msgQueue)
	time.Sleep(100 * time.Millisecond)

	grpcServer.Stop()
	log.Println("Server stopped")
}
