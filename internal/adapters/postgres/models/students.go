package models

type Student struct {
	ID          int `db:"id"`
	ClassroomID int `db:"classroom_id"`
	SchoolID    int `db:"school_id"`
}
