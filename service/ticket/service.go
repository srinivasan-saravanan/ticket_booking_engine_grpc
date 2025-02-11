package ticket

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"ticket_booking_engine/package/ticketpb"
)

type ticketService struct {
	ticketpb.UnimplementedTicketServiceServer
	mu           sync.Mutex
	users        map[string]*ticketpb.Receipt
	sectionSeats map[string][]string
}

func New() *ticketService {
	return &ticketService{
		users: make(map[string]*ticketpb.Receipt),
		sectionSeats: map[string][]string{
			"A": {"A1", "A2", "A3", "A4", "A5"},
			"B": {"B1", "B2", "B3", "B4", "B5"},
		},
	}
}

func (t *ticketService) BookTicket(ctx context.Context, req *ticketpb.BookingRequest) (*ticketpb.Receipt, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	validationErr := t.validateTicketBookingRequest(req)
	if validationErr != nil {
		return nil, validationErr
	}

	section := "B"
	if len(t.sectionSeats["A"]) > len(t.sectionSeats["B"]) {
		section = "A"
	}
	if len(t.sectionSeats[section]) == 0 {
		return nil, fmt.Errorf("no seats available")
	}

	// random seat selection
	seatIndex := rand.Intn(len(t.sectionSeats[section]))
	seat := t.sectionSeats[section][seatIndex]
	t.sectionSeats[section] = append(t.sectionSeats[section][:seatIndex], t.sectionSeats[section][seatIndex+1:]...)

	receipt := &ticketpb.Receipt{
		TicketInfo: &ticketpb.TicketInfo{
			From:  req.TicketInfo.From,
			To:    req.TicketInfo.To,
			User:  req.TicketInfo.User,
			Price: req.TicketInfo.Price,
		},
		SeatSection: section,
		SeatNumber:  seat,
	}

	// Save user
	t.users[req.TicketInfo.User.Email] = receipt

	return receipt, nil
}

func (t *ticketService) GetReceipt(ctx context.Context, req *ticketpb.UserRequest) (*ticketpb.Receipt, error) {
	receipt, exists := t.users[req.GetEmail()]
	if !exists {
		return nil, fmt.Errorf("no receipt not found")
	}
	return receipt, nil
}

func (t *ticketService) GetReceiptsBySeatSection(ctx context.Context, req *ticketpb.SeatSectionRequest) (
	*ticketpb.SectionDetails, error) {

	requestedSection := req.GetSection()
	if requestedSection == "" || !checkKeyInSlice(availableSections, requestedSection) {
		return nil, fmt.Errorf("invalid requested section")
	}

	sectionDetails := &ticketpb.SectionDetails{}
	for _, receipt := range t.users {
		if receipt.SeatSection == requestedSection {
			sectionDetails.UserSeats = append(sectionDetails.UserSeats, &ticketpb.UserSeat{
				User: receipt.TicketInfo.User,
				Seat: receipt.SeatNumber,
			})
		}
	}

	return sectionDetails, nil
}

func (t *ticketService) RemoveUserReceipt(ctx context.Context, req *ticketpb.UserRequest) (*ticketpb.RemoveUserResponse, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	receipt, exists := t.users[req.GetEmail()]
	if !exists {
		return nil, fmt.Errorf("no receipt found")
	}

	// When removing the user, have to free the seat which the user occupied
	t.sectionSeats[receipt.SeatSection] = append(t.sectionSeats[receipt.SeatSection], receipt.SeatNumber)

	delete(t.users, req.GetEmail())
	return &ticketpb.RemoveUserResponse{Success: true}, nil
}

func (t *ticketService) ModifySeat(ctx context.Context, req *ticketpb.ModifySeatRequest) (*ticketpb.Receipt, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if req.Email == "" {
		return nil, fmt.Errorf("invalid email")
	}

	receipt, exists := t.users[req.GetEmail()]
	if !exists {
		return nil, fmt.Errorf("no receipt found")
	}
	log.Printf("Current Receipt: %v", receipt)

	err := t.validateSeatModificationRequest(req, receipt)
	if err != nil {
		return nil, err
	}
	oldSection, oldSeat := receipt.SeatSection, receipt.SeatNumber

	receipt.SeatSection = req.Section
	receipt.SeatNumber = req.SeatNumber

	// Free the Old Seat
	t.sectionSeats[oldSection] = append(t.sectionSeats[oldSection], oldSeat)

	// Occupy Requested Seat
	t.sectionSeats[receipt.SeatSection] = removeElement(t.sectionSeats[receipt.SeatSection], receipt.SeatNumber)

	return receipt, nil
}
