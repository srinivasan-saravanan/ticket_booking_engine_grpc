package ticket

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"ticket_booking_engine/service/ticket/test_cases"
)

var TestTicketService *ticketService

func TestMain(t *testing.M) {
	TestTicketService = New()
	code := t.Run()
	os.Exit(code)
}

func TestTicketService_BookTicket(t *testing.T) {
	for _, tc := range test_cases.BookTicketCases {
		if !tc.Run {
			continue
		}
		t.Run(tc.Case, func(t *testing.T) {
			TestTicketService.users = tc.TicketUsers
			TestTicketService.sectionSeats = tc.AvailableSeatSections
			receipt, err := TestTicketService.BookTicket(context.Background(), tc.Request)
			if tc.ErrorCase {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, receipt.SeatSection)
				assert.NotEmpty(t, receipt.SeatNumber)
				assert.Equal(t, tc.ExpectedSection, receipt.SeatSection)
			}
		})
	}
}

func TestTicketService_GetReceipt(t *testing.T) {
	for _, tc := range test_cases.GetReceiptTestCases {
		if !tc.Run {
			continue
		}
		t.Run(tc.Case, func(t *testing.T) {
			TestTicketService.users = tc.TicketUsers
			receipt, err := TestTicketService.GetReceipt(context.Background(), tc.UserRequest)
			if tc.ErrorCase {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, receipt.SeatSection)
				assert.NotEmpty(t, receipt.SeatNumber)
				assert.Equal(t, receipt.TicketInfo.User.Email, tc.UserRequest.Email)
			}
		})
	}
}

func TestTicketService_GetReceiptsBySeatSection(t *testing.T) {
	for _, tc := range test_cases.GetReceiptsBySectionTestCases {
		if !tc.Run {
			continue
		}
		t.Run(tc.Case, func(t *testing.T) {
			TestTicketService.users = tc.TicketUsers
			sectionDetails, err := TestTicketService.GetReceiptsBySeatSection(context.Background(), tc.SectionRequest)
			if tc.ErrorCase {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, sectionDetails, tc.ExpectedOutput)
			}
		})
	}
}

func TestTicketService_RemoveUserReceipt(t *testing.T) {
	for _, tc := range test_cases.RemoveUserReceiptTestCases {
		if !tc.Run {
			continue
		}
		t.Run(tc.Case, func(t *testing.T) {
			TestTicketService.users = tc.TicketUsers
			TestTicketService.sectionSeats = tc.AvailableSeatSections
			_, err := TestTicketService.RemoveUserReceipt(context.Background(), tc.UserRequest)
			if tc.ErrorCase {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, TestTicketService.users, tc.ExpectedTicketUsers)
				assert.Equal(t, TestTicketService.sectionSeats, tc.ExpectedSeatSections)
			}
		})
	}
}

func TestTicketService_ModifySeat(t *testing.T) {
	for _, tc := range test_cases.ModifySeatRequestCases {
		if !tc.Run {
			continue
		}
		t.Run(tc.Case, func(t *testing.T) {
			TestTicketService.users = tc.TicketUsers
			TestTicketService.sectionSeats = tc.AvailableSeatSections
			receipt, err := TestTicketService.ModifySeat(context.Background(), tc.Request)
			if tc.ErrorCase {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, receipt.SeatNumber)
				assert.NotEmpty(t, receipt.SeatSection)
				assert.Equal(t, receipt.SeatSection, tc.Request.Section)
				assert.Equal(t, receipt.SeatNumber, tc.Request.SeatNumber)
				assert.Equal(t, TestTicketService.sectionSeats, tc.ExpectedAvailableSeatSections)
			}
		})
	}
}
