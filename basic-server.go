package main

import (
	"context"
	"time"

	pb "github.com/r3rivera/r3app-protobuffer-repo/basic-test"

	"log"
	"net"

	"google.golang.org/grpc"
)

type serverHandler struct{}

func (*serverHandler) HealthCheckStatus(ctx context.Context, rqst *pb.HealthCheckStatusRequest) (*pb.HealthCheckStatusResponse, error) {
	log.Println("HealthCheckStatus function is called!")

	status := pb.HealthCheckStatus{
		AppName:       rqst.AppName,
		AppReleaseVer: "1.3.45555",
	}

	resp := pb.HealthCheckStatusResponse{
		HealthStatus: &status,
	}
	return &resp, nil
}

func (*serverHandler) Calculator(ctx context.Context, rqst *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Println("Calculator function is called!", &rqst)

	value := rqst.Payload.NumOne + rqst.Payload.NumTwo
	log.Println("Sum of value is ", value)

	resp := pb.CalculatorResponse{
		Result: value,
	}
	return &resp, nil
}

//Server-Streamin API
func (*serverHandler) NotificationMessage(rqst *pb.NotificationMessageRequest,
	serverStream pb.NotificationMessageService_NotificationMessageServer) error {
	log.Println("Server streaming is in progress!")

	requester := rqst.Requester
	log.Println("Caller of the stream is ", requester)

	notify := pb.NotificationMessage{
		Sender:  "Twitter",
		Message: "Hello There!",
	}

	for i := 0; i < 25; i++ {
		res := &pb.NotificationMessageResponse{
			ResponsePayload: &notify,
		}
		serverStream.Send(res)
		log.Println("Sleeping for 2 seconds... count is ", i)
		time.Sleep(2 * time.Second)
	}

	return nil
}

func main() {
	log.Println("Main Server...")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterHealthCheckStatusServiceServer(server, &serverHandler{})
	pb.RegisterCalculatorServiceServer(server, &serverHandler{})
	pb.RegisterNotificationMessageServiceServer(server, &serverHandler{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
	log.Println("Main server is up and running...")
}
