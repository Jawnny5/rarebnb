package dbrepo

import (
	"errors"
	"log"
	"rarebnb/internal/models"
	"time"
)

func(m *testDBRepo) AllUsers() bool{
	return true
}

//InsertReservation inserts a reservation into the database.
func(m *testDBRepo) InsertReservation(res models.Reservation) (int, error){
	if res.RoomID == 2 {
		return 0, errors.New("unable to create reservation")
	}
	return 1, nil
}

//InsertRoomRestriction inserts a room restriction into the database.
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("failed to insert room restriction")
	}
	return nil
}

//SearchAvailabilityByDates returns true if availability exists for roomID, false if not
func (m *testDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID int)(bool, error){
		// Set up test time
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	// Test to fail the query -- specify 2060-01-01 as start
	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return false, errors.New("some error")
	}

	// If the start date is after 2049-12-31, then return false,
	// indicating no availability;
	if start.After(t) {
		return false, nil
	}

	// Otherwise, we have availability
	return true, nil
}

//SearchAvailabilityForAllRooms returns a slice of rooms, if any, for given date range 
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time)([]models.Room, error){
	var rooms []models.Room

	return rooms, nil
}

//GetRoomById gets a room by ID.
func (m *testDBRepo) GetRoomById(id int) (models.Room, error){
	var room models.Room
	if id > 2 {
		return room, errors.New("Error")
	}
	return room, nil
}

func (m *testDBRepo) GetUserById(id int) (models.User, error){
	var u models.User
	
	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil 
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error){
	return 1, "", nil
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error){
	var reservations []models.Reservation
	return reservations, nil
}

//AllNewReservations returns a slice of all reservations.
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error){
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDBRepo) GetReservationById(id int)(models.Reservation, error){
	var res models.Reservation
	return res, nil
}


func (m *testDBRepo) UpdateReservation(u models.Reservation) error{
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

//UpdatedProcessedForReservation updates processed for a reservation by id
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error ){
	var rooms []models.Room 
	return rooms, nil
}


