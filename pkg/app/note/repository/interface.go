package repository

import (
	"note-manager/pkg/domain/note"
)

// Repository is repository
type Repository interface {
	GetNotes(string, int) ([]note.Note, error)
	AddNotes(notes []note.Note) error
	UpdateNote(note.Note) error
	DeleteNote(id string) error
}
