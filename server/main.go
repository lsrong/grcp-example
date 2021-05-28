package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lsrong/grpc-alarm/alarm"
	"google.golang.org/grpc"
)

const (
	port = ":10050"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAlarmServiceServer
}

func (s *server) SetAlarmMessage(_ context.Context, in *pb.AlarmRequest) (*pb.AlarmReply, error) {

	log.Println(in.Type, in.FileName, in.DownloadUrl, in.Md5Value, in.SzUserName, in.SzPassword, in.SzIP, in.NPort)

	return &pb.AlarmReply{Message: "success", Code: pb.ResponseCode_Ok_200}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAlarmServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
