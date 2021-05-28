package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lsrong/grpc-alarm/alarm"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10050"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAlarmServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/**
		SzIP        string     `protobuf:"bytes,2,opt,name=szIP,proto3" json:"szIP,omitempty"`
	NPort       uint32     `protobuf:"varint,3,opt,name=nPort,proto3" json:"nPort,omitempty"`
	SzUserName  string     `protobuf:"bytes,4,opt,name=szUserName,proto3" json:"szUserName,omitempty"`
	SzPassword  string     `protobuf:"bytes,5,opt,name=szPassword,proto3" json:"szPassword,omitempty"`
	DownloadUrl string     `protobuf:"bytes,6,opt,name=downloadUrl,proto3" json:"downloadUrl,omitempty"`
	FileName    string     `protobuf:"bytes,7,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Md5Value    string     `protobuf:"bytes,8,opt,name=md5Value,proto3" json:"md5Value,omitempty"`
	*/
	request := &pb.AlarmRequest{
		Type:        pb.CameraType_DaHua,
		SzIP:        "127.0.0.1",
		NPort:       1000,
		SzUserName:  "admin",
		SzPassword:  "admin",
		DownloadUrl: "http://localhost",
		FileName:    "filename",
		Md5Value:    "test",
	}

	r, err := c.SetAlarmMessage(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: code=%d, message=%s", r.GetCode(), r.GetMessage())
}
