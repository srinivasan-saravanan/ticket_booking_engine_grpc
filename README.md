# Ticket Booking Engine Using gRPC

## Overview

This project is used to book a train ticket for a single route using Go and gRPC. It does not have 
database layer. The booking data stored in memory for a particular session.

## Prerequisites

- Go 1.20

## Project Structure
The following is the directory structure of the project:

```plaintext
ticket_booking_engine_grpc/
│
├── README.md                                         # Project description and setup instructions
├── go.mod                                            # Go module dependencies
├── go.sum                                            # Go module checksums
├── main.go                                           # Main server entry point
├── app
│   ├── app.go                                        # File which started gRPC server and listens for communication              
├── service/
│   └── ticket/ 
│       └── test_cases/
│           ├── book_ticket_testcases.go              # Book Ticket Testcases
│           ├── get_receipt_testcases.go              # Get Receipt Testcases
│           ├── get_receipts_by_section_testcases.go  # Get Receipts by Section Testcases 
│           ├── modify_seat_testcases.go              # Modify Seat Testcases
│           ├── remove_user_receipt_testcases.go      # Remove user booking Testcases  
│   ├── service.go                                    # Ticket service implementation
│   ├── helpers.go                                    # Helpers for Ticket service implementation
│   ├── ticket_test.go                                # Unit Testing for Ticket Service Implementaion
├── proto/
│   ├── ticket.proto                                  # gRPC service definition for ticket booking
│   ├── user.proto                                    # gRPC service definition for user management
├── package/
│   └── ticketpb/                                     # Generated Go code for ticket service
│   └── userpb/                                       # Generated Go code for user service
├── client.go                                         # Client to interact with gRPC API
├── generate.sh                                       # Shell script to generate gRPC service definition
```


## Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/your-username/ticket_booking_engine_grpc.git
    cd ticket_booking_engine_grpc
    ```

2. **Download Dependencies**:
   ```bash
    go mod tidy
    ```

3. **Start gRPC server**:
   ```bash
   go run main.go
   ```
   
4. **To Run gRPC client**:
   ```bash
   go run client.go
   ```

5. **To Run Unit Testing**:
   ```bash
   go test -v ./service/ticket -v
   ```   
   



