package repository

import (
	"rarebnb/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)  
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int)(bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time)([]models.Room, error)
	GetRoomById(id int) (models.Room, error)
	GetUserById(id int) (models.User, error)
	Authenticate(email, testPassword string) (int, string, error)
	AllReservations() ([]models.Reservation, error)
}
