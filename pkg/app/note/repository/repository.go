package repository

import (
	"note-manager/pkg/domain/note"
	"note-manager/pkg/infra/db"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
)

type noteRepository struct{}

// NewNoteRepository new a repository
func NewNoteRepository() Repository {
	client = db.NewClient()
	return &noteRepository{}
}

func (u *noteRepository) GetNotes() ([]note.Note, error) {
	filter := bson.D{}
	var notes []note.Note
	ctx := context.Background()
	collection := client.Database("note").Collection("notes")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
