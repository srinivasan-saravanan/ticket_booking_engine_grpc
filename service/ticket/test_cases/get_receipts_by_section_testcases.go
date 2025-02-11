package test_cases

import (
	"fmt"
	"ticket_booking_engine/package/ticketpb"
	"ticket_booking_engine/package/userpb"
)

type GetReceiptsBySectionTestCase struct {
	Case           string
	SectionRequest *ticketpb.SeatSectionRequest
	TicketUsers    map[string]*ticketpb.Receipt
	ErrorCase      bool
	ExpectedErr    error
	ExpectedOutput *ticketpb.SectionDetails
	Run            bool
}

var GetReceiptsBySectionTestCases = []GetReceiptsBySectionTestCase{
	{
		Case:           "Empty Section Provided",
		SectionRequest: &ticketpb.SeatSectionRequest{Section: ""},
		TicketUsers:    map[string]*ticketpb.Receipt{},
		ErrorCase:      true,
		ExpectedErr:    fmt.Errorf("invalid requested section"),
		Run:            true,
	},
	{
		Case:           "Invalid Section Provided",
		SectionRequest: &ticketpb.SeatSectionRequest{Section: "C"},
		TicketUsers:    map[string]*ticketpb.Receipt{},
		ErrorCase:      true,
		ExpectedErr:    fmt.Errorf("invalid requested section"),
		Run:            true,
	},
	{
		Case:           "Valid Case",
		SectionRequest: &ticketpb.SeatSectionRequest{Section: "A"},
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
			"test2@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test2@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "A", SeatNumber: "A2",
			},
			"test3@gmail.com": {
				TicketInfo: &ticketpb.TicketInfo{
					From: "London", To: "France", Price: 20.0,
					User: &userpb.User{
						Email: "test2@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
				},
				SeatSection: "B", SeatNumber: "B1",
			},
		},
		ErrorCase: false,
		ExpectedOutput: &ticketpb.SectionDetails{
			UserSeats: []*ticketpb.UserSeat{
				{
					User: &userpb.User{
						Email: "test1@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
					Seat: "A1",
				},
				{
					User: &userpb.User{
						Email: "test2@gmail.com", FirstName: "First Name", LastName: "Last Name",
					},
					Seat: "A2",
				},
			},
		},
		ExpectedErr: nil,
		Run:         true,
	},
}
