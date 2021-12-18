package repository

import (
	"fmt"
	"note-manager/pkg/domain/note"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/db"
	"note-manager/pkg/infra/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_noteRepository_GetNotes(t *testing.T) {
	config.Init()
	db.Init(logger.NewMockLogger())
	repo := NewNoteRepository()
	ns, err := repo.GetNotes("李維", "", 1)
	assert.NotNil(t, ns)
	assert.NoError(t, err)
}

func Test_noteRepository_GetNotes_Tags(t *testing.T) {
	config.Init()
	db.Init(logger.NewMockLogger())
	repo := NewNoteRepository()
	ns, err := repo.GetNotes("李維", "人類學", 1)
	assert.NotNil(t, ns)
	assert.NoError(t, err)
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

func Test_noteRepository_GetTags(t *testing.T) {
	config.Init()
	db.Init(logger.NewMockLogger())
	repo := NewNoteRepository()
	tags, _ := repo.GetTags()
	fmt.Println(tags)
}
