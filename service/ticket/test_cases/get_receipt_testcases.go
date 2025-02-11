package test_cases

import (
	"fmt"
	"ticket_booking_engine/package/ticketpb"
	"ticket_booking_engine/package/userpb"
)

type GetReceiptTestCase struct {
	Case            string
	UserRequest     *ticketpb.UserRequest
	TicketUsers     map[string]*ticketpb.Receipt
	ExpectedReceipt *ticketpb.Receipt
	ErrorCase       bool
	ExpectedErr     error
	Run             bool
}

var GetReceiptTestCases = []GetReceiptTestCase{
	{
		Case:        "Receipt Not Found",
		UserRequest: &ticketpb.UserRequest{Email: "test@gmail.com"},
		TicketUsers: map[string]*ticketpb.Receipt{
			"test1@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A1",
			},
		},
		ExpectedReceipt: nil,
		ErrorCase:       true,
		ExpectedErr:     fmt.Errorf("no receipt not found"),
		Run:             true,
	},
	{
		Case:        "Valid Case",
		UserRequest: &ticketpb.UserRequest{Email: "test@gmail.com"},
		TicketUsers: map[string]*ticketpb.Receipt{
			"test@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A1",
			},
		},
		ExpectedReceipt: &ticketpb.Receipt{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
			SeatSection: "A", SeatNumber: "A1",
		},
		ErrorCase:   false,
		ExpectedErr: nil,
		Run:         true,
	},
}
