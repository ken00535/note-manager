package usecase

import (
	"note-manager/pkg/app/note/repository"
	"note-manager/pkg/domain/note"
)

type noteUsecase struct {
	repo repository.Repository
}

// NewNoteUsecase will create new usecase
func NewNoteUsecase(repo repository.Repository) Usecase {
	return &noteUsecase{
		repo: repo,
	}
}

func (u *noteUsecase) GetNotes() ([]note.Note, error) {
	return u.repo.GetNotes()
}

func (u *noteUsecase) AddNotes(notes []note.Note) error {
	return u.repo.AddNotes(notes)
}

func (u *noteUsecase) DeleteNote(id string) error {
	return u.repo.DeleteNote(id)
}
