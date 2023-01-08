package repository

import "github.com/victoorraphael/coordinator/internal/domain"

type IStudent interface {
	Get() []domain.Student
}
