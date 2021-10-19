package usecase

import "note-manager/pkg/domain/note"

// Usecase is usecase
type Usecase interface {
	GetNotes() ([]note.Note, error)
}
