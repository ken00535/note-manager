package repository

import (
	"note-manager/pkg/app/note/entity"
)

// Repository is repository
type Repository interface {
	GetNotes(string, string, int) ([]entity.Note, error)
	AddNotes(notes []entity.Note) ([]string, error)
	UpdateNote(entity.Note) error
	DeleteNote(id string) error
	GetTags() ([]entity.Tag, error)
}
