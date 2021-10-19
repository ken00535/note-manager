package repository

import (
	"note-manager/pkg/domain/note"
)

// Repository is repository
type Repository interface {
	GetNotes() ([]note.Note, error)
}
