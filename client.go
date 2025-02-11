package main

import (
	"context"
	"log"
	"ticket_booking_engine/package/userpb"
	"time"

	"google.golang.org/grpc"
	"ticket_booking_engine/package/ticketpb"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connet: %v", err)
	}
	defer conn.Close()

	client := ticketpb.NewTicketServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Book Ticket Block Starts
	ticket, err := client.BookTicket(ctx, &ticketpb.BookingRequest{
		TicketInfo: &ticketpb.TicketInfo{
			From: "London", To: "France", Price: 20.0,
			User: &userpb.User{
				FirstName: "Srinivasan",
				LastName:  "Saravanan",
				Email:     "srinivasan@gmail.com",
			},
		},
	})

	if err != nil {
		log.Fatalf("could not book ticket: %v", err.Error())
	}
	log.Printf("Booked Receipt: %v", ticket)
	// Book Ticket Block Ends

	// Get Receipt Block Starts
	receipt, err := client.GetReceipt(ctx, &ticketpb.UserRequest{Email: "srinivasan@gmail.com"})
	if err != nil {
		log.Fatalf("could not get receipt: %v", err.Error())
	}
	log.Printf("Ticket Receipt: %v", receipt)
	// Get Receipt Block Ends

	// Get Receipts By Section Block Starts
	section, err := client.GetReceiptsBySeatSection(ctx, &ticketpb.SeatSectionRequest{Section: receipt.GetSeatSection()})
	if err != nil {
		log.Fatalf("could not get receipts: %v", err.Error())
	}
	log.Printf("Receipts By Section: %v", section)
	// Get Receipts By Section Block Ends

	// Modify Seat Block Starts
	seat, err := client.ModifySeat(ctx, &ticketpb.ModifySeatRequest{Email: "srinivasan@gmail.com", Section: "A", SeatNumber: "A5"})
	if err != nil {
		log.Fatalf("could not modify seat: %v", err.Error())
	}
	log.Printf("Ticket Receipt after Seat Modification: %v", seat)
	// Modify Seat Block Ends

	// Get Users By Section After Seat Modification Starts
	section, err = client.GetReceiptsBySeatSection(ctx, &ticketpb.SeatSectionRequest{Section: seat.GetSeatSection()})
	if err != nil {
		log.Fatalf("could not get receipts: %v", err.Error())
	}
	log.Printf("Receipts By Section, After Seat Modification: %v", section)
	// Get Users By Section After Seat Modification Ends

	// Remove User From List Block Starts
	status, err := client.RemoveUserReceipt(ctx, &ticketpb.UserRequest{Email: receipt.TicketInfo.User.Email})
	if err != nil {
		log.Fatalf("could not remove receipt: %v", err.Error())
	}
	log.Printf("Remove User Request: %v", status)
	// Remove User From List Block Ends
}
