package models

type Subject struct {
	ID   int    `db:"id"`
	Nome string `db:"nome"`
}

type SubjectClassroom struct {
	SubjectID   int `db:"subject_id"`
	ClassroomID int `db:"classroom_id"`
}
