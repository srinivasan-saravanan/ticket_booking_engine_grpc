package test_cases

import (
	"fmt"
	"ticket_booking_engine/package/ticketpb"
	"ticket_booking_engine/package/userpb"
)

type RemoveUserReceiptTestCase struct {
	Case                  string
	UserRequest           *ticketpb.UserRequest
	TicketUsers           map[string]*ticketpb.Receipt
	ExpectedTicketUsers   map[string]*ticketpb.Receipt
	AvailableSeatSections map[string][]string
	ExpectedSeatSections  map[string][]string
	ErrorCase             bool
	ExpectedErr           error
	Run                   bool
}

var RemoveUserReceiptTestCases = []RemoveUserReceiptTestCase{
	{
		Case:        "Receipt not Found",
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
		AvailableSeatSections: map[string][]string{
			"A": {"A2", "A3", "A4", "A5"},
			"B": {"B1", "B2", "B3", "B4", "B5"},
		},
		ExpectedSeatSections: map[string][]string{
			"A": {"A2", "A3", "A4", "A5"},
			"B": {"B1", "B2", "B3", "B4", "B5"},
		},
		ErrorCase:   true,
		ExpectedErr: fmt.Errorf("no receipt found"),
		Run:         true,
	},
	{
		Case:        "Valid Case",
		UserRequest: &ticketpb.UserRequest{Email: "test@gmail.com"},
		TicketUsers: map[string]*ticketpb.Receipt{
			"test@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A1",
			},
		},
		ExpectedTicketUsers: map[string]*ticketpb.Receipt{},
		AvailableSeatSections: map[string][]string{
			"A": {"A2", "A3", "A4", "A5"},
			"B": {"B1", "B2", "B3", "B4", "B5"},
		},
		ExpectedSeatSections: map[string][]string{
			"A": {"A2", "A3", "A4", "A5", "A1"},
			"B": {"B1", "B2", "B3", "B4", "B5"},
		},
		ErrorCase:   false,
		ExpectedErr: nil,
		Run:         true,
	},
}
