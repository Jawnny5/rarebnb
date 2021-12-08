package models

import (
	"time"
)

//User defines the users model
type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	Password string
	AccessLevel int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Room defines the room model
type Room struct {
	ID int
	RoomName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Restriction defines the restriction model
type Restriction struct {
	ID int 
	RestrictionName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Reservation defines the reservation model
type Reservation struct {
	ID int 
	RestrictionName string
	FirstName string
	LastName string
	Email string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Room
}

type RoomRestriction struct {
	ID int 
	RoomID int
	ReservationID int
	RestrictionID int
	StartDate time.Time
	EndDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Room
	Reservation Reservation
	Restriction Restriction
}

type MailData struct {
	To string
	From string
	Subject string
	Content string
	Template string
}