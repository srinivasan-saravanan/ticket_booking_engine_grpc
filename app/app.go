package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"ticket_booking_engine/package/ticketpb"
	"ticket_booking_engine/service/ticket"
)

func registerServices(server *grpc.Server) {
	ticketServer := ticket.New()
	ticketpb.RegisterTicketServiceServer(server, ticketServer)
}

func Start() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	registerServices(server)

	reflection.Register(server)
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
