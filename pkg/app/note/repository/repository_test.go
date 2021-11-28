package repository

import (
	"note-manager/pkg/domain/note"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/db"
	"note-manager/pkg/infra/logger"
	"testing"
)

func Test_noteRepository_GetNotes(t *testing.T) {
	config.Init()
	db.Init(logger.NewMockLogger())
	repo := NewNoteRepository()
	repo.GetNotes("李維", 1)
}

func Test_noteRepository_AddNotes(t *testing.T) {
	config.Init()
	db.Init(logger.NewMockLogger())
	repo := NewNoteRepository()
	notes := []note.Note{
		{
			Content: "test",
			Comment: "test",
		},
	}
	repo.AddNotes(notes)
}
