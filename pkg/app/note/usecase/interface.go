package usecase

import (
	"note-manager/pkg/app/note/entity"
)

// Usecase is usecase
type Usecase interface {
	GetNotes(string, string, int) ([]entity.Note, error)
	AddNotes(notes []entity.Note) ([]string, error)
	UpdateNote(entity.Note) error
	DeleteNote(string) error
	GetTags() ([]entity.Tag, error)
}
