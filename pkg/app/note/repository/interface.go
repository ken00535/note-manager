package repository

import (
	"note-manager/pkg/domain/note"
)

// Repository is repository
type Repository interface {
	GetNotes(string, string, int) ([]note.Note, error)
	AddNotes(notes []note.Note) ([]string, error)
	UpdateNote(note.Note) error
	DeleteNote(id string) error
}
