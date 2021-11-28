package repository

import (
	"note-manager/pkg/domain/note"
	"note-manager/pkg/infra/db"
	"note-manager/pkg/infra/logger"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	log    logger.Logger
	client *mongo.Client
)

type noteRepository struct{}

// NewNoteRepository new a repository
func NewNoteRepository() Repository {
	log = logger.New()
	client = db.NewClient()
	return &noteRepository{}
}

func (u *noteRepository) GetNotes(kw string, page int) ([]note.Note, error) {
	options := options.Find()
	type Document struct {
		ID      primitive.ObjectID `bson:"_id"`
		Content string             `json:"content"`
		Comment string             `json:"comment"`
	}
	var docs []Document
	var notes []note.Note
	ctx := context.Background()
	collection := client.Database("note").Collection("notes")
	log.Info("start: ", time.Now())
	cursor, err := collection.Find(
		ctx,
		bson.M{"content": primitive.Regex{Pattern: kw, Options: ""}},
		options.SetLimit(10),
		options.SetSkip(int64(10*(page-1))),
	)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	log.Info("end: ", time.Now())
	for _, d := range docs {
		notes = append(notes, note.Note{
			ID:      d.ID.Hex(),
			Content: d.Content,
			Comment: d.Comment,
		})
	}
	return notes, nil
}

func (u *noteRepository) AddNotes(notes []note.Note) ([]string, error) {
	var res []string
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
	result, err := collection.InsertMany(ctx, ds)
	if err != nil {
		log.Error(err)
		return res, err
	}
	for _, id := range result.InsertedIDs {
		oid := id.(primitive.ObjectID)
		res = append(res, oid.Hex())
	}
	return res, nil
}

func (u *noteRepository) UpdateNote(n note.Note) error {
	type Document struct {
		ID      primitive.ObjectID `bson:"_id"`
		Content string             `json:"content"`
		Comment string             `json:"comment"`
	}
	ctx := context.Background()
	idPrimitive, err := primitive.ObjectIDFromHex(n.ID)
	d := Document{
		ID:      idPrimitive,
		Content: n.Content,
		Comment: n.Comment,
	}
	collection := client.Database("note").Collection("notes")
	_, err = collection.UpdateOne(ctx, bson.M{"_id": idPrimitive}, bson.M{"$set": d})
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return err
	}
	return nil
}
