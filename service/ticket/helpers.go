package ticket

import (
	"fmt"
	"ticket_booking_engine/package/ticketpb"
)

var availableSections = []string{"A", "B"}

func (t *ticketService) validateTicketBookingRequest(req *ticketpb.BookingRequest) error {
	if req.TicketInfo.User == nil {
		return fmt.Errorf("invalid user info")
	}

	if req.TicketInfo.User.Email == "" {
		return fmt.Errorf("invalid user email")
	}

	if _, exists := t.users[req.TicketInfo.User.Email]; exists {
		return fmt.Errorf("user with email %s already exists", req.TicketInfo.User.Email)
	}

	return nil
}

func (t *ticketService) validateSeatModificationRequest(req *ticketpb.ModifySeatRequest, currentReceipt *ticketpb.Receipt) error {
	// Returns error if user already occupied the same seat
	if !checkKeyInSlice(availableSections, req.Section) {
		return fmt.Errorf("invalid section")
	}

	if currentReceipt.SeatSection == req.Section && currentReceipt.SeatNumber == req.SeatNumber {
		return fmt.Errorf("invalid seat number")
	}

	if len(t.sectionSeats[req.Section]) == 0 {
		return fmt.Errorf("no seats available in given section, could not proceed with seat modification request")
	}

	availableSeats := t.sectionSeats[req.Section]
	if !checkKeyInSlice(availableSeats, req.SeatNumber) {
		return fmt.Errorf("invalid seat number, the requested seat is already taken")
	}

	//for _, receipt := range t.users {
	//	if receipt.SeatSection == req.Section && receipt.SeatNumber == req.SeatNumber {
	//		return fmt.Errorf("invalid seat number, the requested seat is already taken")
	//	}
	//}

	return nil
}

func checkKeyInSlice(items []string, key string) bool {
	for _, item := range items {
		if item == key {
			return true
		}
	}
	return false
}

func removeElement(slice []string, element string) []string {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
