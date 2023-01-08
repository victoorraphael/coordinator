package models

type Professor struct {
	ID       int `db:"id"`
	SchoolID int `db:"school_id"`
}

type ProfessorSubject struct {
	ProfessorID int `db:"professor_id"`
	SubjectID   int `db:"subject_id"`
}
