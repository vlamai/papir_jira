package main

import (
	"github.com/vlamai/papir_jira/config"
	"github.com/vlamai/papir_jira/jira"
	"github.com/vlamai/papir_jira/proto/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	jiraTracker, err := jira.New(config.NewJira())
	if err != nil {
		log.Fatalf("init jira | %v", err)
	}
	lis, err := net.Listen("tcp", config.NewGRPC().ConnectionSting)
	if err != nil {
		log.Fatalf("failed to listen | %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskTrackerServer(s, jiraTracker)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve | %v", err)
	}
}
