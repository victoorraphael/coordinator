package models

type School struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	AddressID int    `db:"address_id"`
}

type SchoolClassroom struct {
	SchoolID    int `db:"school_id"`
	ClassroomID int `db:"classroom_id"`
}
