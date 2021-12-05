package usecase

import "note-manager/pkg/domain/note"

// Usecase is usecase
type Usecase interface {
	GetNotes(string, string, int) ([]note.Note, error)
	AddNotes(notes []note.Note) ([]string, error)
	UpdateNote(note.Note) error
	DeleteNote(string) error
}
