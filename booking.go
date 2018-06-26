package main

type Booking struct {
	Court		string	`json:"court"`
	Days		string	`json:"days"`
	Hour		string	`json;"hour"`
	Min		string	`json:"min"`
	Timeslot	string	`json:"timeslot"`
	PlayerA		string	`json:"playerA"`
	PlayerB		string	`json:"playerB"`
	BookingLink	string	`json:"bookingLink"`
	Booked		bool	`json:"booked"`
}

type Bookings []Booking
