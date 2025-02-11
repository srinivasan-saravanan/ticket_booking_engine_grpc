package test_cases

import (
	"fmt"
	"ticket_booking_engine/package/ticketpb"
	"ticket_booking_engine/package/userpb"
)

type BookTicket struct {
	Case                  string
	Request               *ticketpb.BookingRequest
	TicketUsers           map[string]*ticketpb.Receipt
	AvailableSeatSections map[string][]string
	ExpectedSection       string
	ErrorCase             bool
	ExpectedErr           error
	Run                   bool
}

var defaultAvailableSections = map[string][]string{
	"A": {"A1", "A2", "A3", "A4", "A5"},
	"B": {"B1", "B2", "B3", "B4", "B5"},
}

var BookTicketCases = []BookTicket{
	{
		Case: "Invalid user Object",
		Request: &ticketpb.BookingRequest{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
			},
		},
		TicketUsers:           make(map[string]*ticketpb.Receipt),
		AvailableSeatSections: defaultAvailableSections,
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("invalid user info"),
		Run:                   true,
	},
	{
		Case: "Invalid user email",
		Request: &ticketpb.BookingRequest{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "", FirstName: "First Name", LastName: "Last Name",
				},
			},
		},
		TicketUsers:           make(map[string]*ticketpb.Receipt),
		AvailableSeatSections: defaultAvailableSections,
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("invalid user info"),
		Run:                   true,
	},
	{
		Case: "User already has ticket",
		Request: &ticketpb.BookingRequest{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
		},
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
		AvailableSeatSections: defaultAvailableSections,
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("user with email test@gmail.com already exists"),
		Run:                   true,
	},
	{
		Case: "Valid Ticket Booking",
		Request: &ticketpb.BookingRequest{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
		},
		TicketUsers:           make(map[string]*ticketpb.Receipt),
		AvailableSeatSections: defaultAvailableSections,
		ExpectedSection:       "B",
		ErrorCase:             false,
		ExpectedErr:           nil,
		Run:                   true,
	},
	{
		Case: "Picking the Section which has more seats",
		Request: &ticketpb.BookingRequest{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
		},
		TicketUsers: make(map[string]*ticketpb.Receipt),
		AvailableSeatSections: map[string][]string{
			"A": {"A1", "A2", "A3", "A4", "A5"},
			"B": {"B1", "B2", "B3", "B4"},
		},
		ExpectedSection: "A",
		ErrorCase:       false,
		ExpectedErr:     nil,
		Run:             true,
	},
	{
		Case: "Both Section has 0 Seats",
		Request: &ticketpb.BookingRequest{
			TicketInfo: &ticketpb.TicketInfo{
				From: "London", To: "France", Price: 20.0,
				User: &userpb.User{
					Email: "test@gmail.com", FirstName: "First Name", LastName: "Last Name",
				},
			},
		},
		TicketUsers:           make(map[string]*ticketpb.Receipt),
		AvailableSeatSections: map[string][]string{"A": {}, "B": {}},
		ErrorCase:             true,
		ExpectedErr:           fmt.Errorf("no seats available"),
		Run:                   true,
	},
}
