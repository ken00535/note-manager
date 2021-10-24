package usecase

import "note-manager/pkg/domain/note"

// Usecase is usecase
type Usecase interface {
	GetNotes(string, int) ([]note.Note, error)
	AddNotes(notes []note.Note) error
	UpdateNote(note.Note) error
	DeleteNote(string) error
}
