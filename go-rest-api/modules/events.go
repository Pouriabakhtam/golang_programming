package modules

import (
	"errors"
	"log"
	"time"

	"restapi.ca/db"
)

type Events struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Events{}

func (e *Events) Save() error {
	query := `
	INSERT INTO events(name, description, location, datetime, user_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("Couldn't prepare inserting")
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return errors.New("Couldn't insert any DATA")
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
	// events = append(events, *e)
}

func GetAllEvents() ([]Events, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allevents []Events

	for rows.Next() {
		var event Events
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		allevents = append(allevents, event)
	}
	return allevents, nil
}

func GetEventbyID(id int64) (*Events, error) {
	query := `
	SELECT * FROM EVENTS WHERE id = ?
	`
	// QueryRow gives us one row
	row := db.DB.QueryRow(query, id)
	var event Events
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Events) UpdateEvent() error {
	query := `
	UPDATE EVENTS
	SET name = ?, description =?, location = ?, datetime = ?
	WHERE id = ? 
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Events) DeleteEvent() error {
	query := `
	DELETE FROM EVENTS WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}
