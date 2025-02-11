package test_cases

import (
	"fmt"
	"ticket_booking_engine/package/ticketpb"
	"ticket_booking_engine/package/userpb"
)

type ModifySeatRequest struct {
	Case                          string
	Request                       *ticketpb.ModifySeatRequest
	TicketUsers                   map[string]*ticketpb.Receipt
	AvailableSeatSections         map[string][]string
	ExpectedAvailableSeatSections map[string][]string
	ExpectedReceipt               *ticketpb.Receipt
	ErrorCase                     bool
	ExpectedErr                   error
	Run                           bool
}

var ModifySeatRequestCases = []ModifySeatRequest{
	{
		Case:                  "Empty Email",
		Request:               &ticketpb.ModifySeatRequest{Email: ""},
		TicketUsers:           map[string]*ticketpb.Receipt{},
		AvailableSeatSections: map[string][]string{},
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("invalid email"),
		Run:                   true,
	},
	{
		Case:                  "Email Not Exists in Receipts",
		Request:               &ticketpb.ModifySeatRequest{Email: "test@gmail.com"},
		TicketUsers:           map[string]*ticketpb.Receipt{},
		AvailableSeatSections: map[string][]string{},
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("no receipt found"),
		Run:                   true,
	},
	{
		Case:    "Invalid Ticket Section",
		Request: &ticketpb.ModifySeatRequest{Email: "test@gmail.com", Section: "C"},
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
		AvailableSeatSections: map[string][]string{},
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("invalid section"),
		Run:                   true,
	},
	{
		Case:    "Requested for Same Seat",
		Request: &ticketpb.ModifySeatRequest{Email: "test@gmail.com", Section: "A", SeatNumber: "A1"},
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
		AvailableSeatSections: map[string][]string{},
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("invalid seat number"),
		Run:                   true,
	},
	{
		Case:    "Requested Section Has 0 Seats",
		Request: &ticketpb.ModifySeatRequest{Email: "test@gmail.com", Section: "B", SeatNumber: "B1"},
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
		AvailableSeatSections: map[string][]string{
			"A": {}, "B": {},
		},
		ErrorCase:   true,
		ExpectedErr: fmt.Errorf("no seats available in given section, could not proceed with seat modification request"),
		Run:         true,
	},
	{
		Case:    "Requested Seat is already Taken",
		Request: &ticketpb.ModifySeatRequest{Email: "test@gmail.com", Section: "A", SeatNumber: "A2"},
		TicketUsers: map[string]*ticketpb.Receipt{
			"test@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "B", SeatNumber: "B1",
			},
			"test1@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A2",
			},
		},
		AvailableSeatSections: map[string][]string{
			"A": {"A3"}, "B": {"B3"},
		},
		ErrorCase:   true,
		ExpectedErr: fmt.Errorf("invalid seat number, the requested seat is already taken"),
		Run:         true,
	},
	{
		Case:    "Valid Case",
		Request: &ticketpb.ModifySeatRequest{Email: "test@gmail.com", Section: "A", SeatNumber: "A2"},
		TicketUsers: map[string]*ticketpb.Receipt{
			"test@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "B", SeatNumber: "B1",
			},
			"test1@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A2",
			},
		},
		AvailableSeatSections: map[string][]string{
			"A": {"A2", "A3"}, "B": {"B3"},
		},
		ExpectedAvailableSeatSections: map[string][]string{
			"A": {"A3"}, "B": {"B3", "B1"},
		},
		ExpectedReceipt: &ticketpb.Receipt{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
			SeatSection: "A", SeatNumber: "A2",
		},
		ErrorCase:   false,
		ExpectedErr: nil,
		Run:         true,
	},
	{
		Case:    "Valid Case 2",
		Request: &ticketpb.ModifySeatRequest{Email: "test@gmail.com", Section: "B", SeatNumber: "B5"},
		TicketUsers: map[string]*ticketpb.Receipt{
			"test@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "B", SeatNumber: "B1",
			},
			"test1@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A2",
			},
		},
		AvailableSeatSections: map[string][]string{
			"A": {"A2", "A3"}, "B": {"B5"},
		},
		ExpectedAvailableSeatSections: map[string][]string{
			"A": {"A2", "A3"}, "B": {"B1"},
		},
		ExpectedReceipt: &ticketpb.Receipt{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
			SeatSection: "B", SeatNumber: "B5",
		},
		ErrorCase:   false,
		ExpectedErr: nil,
		Run:         true,
	},
}
