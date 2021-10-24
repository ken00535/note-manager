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

func (u *noteUsecase) GetNotes(kw string) ([]note.Note, error) {
	return u.repo.GetNotes(kw)
}

func (u *noteUsecase) AddNotes(notes []note.Note) error {
	return u.repo.AddNotes(notes)
}

func (u *noteUsecase) UpdateNote(n note.Note) error {
	return u.repo.UpdateNote(n)
}

func (u *noteUsecase) DeleteNote(id string) error {
	return u.repo.DeleteNote(id)
}
