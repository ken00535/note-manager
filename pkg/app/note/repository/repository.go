package repository

import (
	"fmt"
	"note-manager/pkg/domain/note"
	"note-manager/pkg/infra/db"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	type Document struct {
		ID      primitive.ObjectID `bson:"_id"`
		Content string             `json:"content"`
		Comment string             `json:"comment"`
	}
	var docs []Document
	var notes []note.Note
	ctx := context.Background()
	collection := client.Database("note").Collection("notes")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	for _, d := range docs {
		notes = append(notes, note.Note{
			ID:      d.ID.Hex(),
			Content: d.Content,
			Comment: d.Comment,
		})
	}
	return notes, nil
}

func (u *noteRepository) AddNotes(notes []note.Note) error {
	type Document struct {
		Content string `json:"content"`
		Comment string `json:"comment"`
	}
	var ds []interface{}
	for _, n := range notes {
		ds = append(ds, Document{
			Content: n.Content,
			Comment: n.Comment,
		})
	}
	ctx := context.Background()
	collection := client.Database("note").Collection("notes")
	_, err := collection.InsertMany(ctx, ds)
	if err != nil {
		return err
	}
	return nil
}

func (u *noteRepository) DeleteNote(id string) error {
	ctx := context.Background()
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	collection := client.Database("note").Collection("notes")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		fmt.Println()
		return err
	}
	return nil
}
