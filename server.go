package main

import (
	"context"
	"github.com/zh1gr/go-grpc-domain/invoice"
	"google.golang.org/grpc"
	"log"
	"net"
)

type invoiceServer struct {
	invoice.UnimplementedInvoiceServer
}

func (s invoiceServer) Create(ctx context.Context, req *invoice.CreateRequest) (*invoice.CreateResponse, error) {
	return &invoice.CreateResponse{
		Pdf:  []byte("hello pdf"),
		Docx: []byte("hello docx"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	service := &invoiceServer{}

	invoice.RegisterInvoiceServer(grpcServer, service)

	log.Println("Listening on :8080")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
