package repository

import "github.com/victoorraphael/coordinator/internal/domain"

type IProfessor interface {
	Get() []domain.Professor
}