package repository

import (
	"note-manager/pkg/app/note/entity"
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

func (u *noteRepository) GetNotes(kw string, tag string, page int) ([]note.Note, error) {
	options := options.Find()
	type Document struct {
		ID        primitive.ObjectID `bson:"_id"`
		Content   string             `bson:"content"`
		Comment   string             `bson:"comment"`
		CreatedAt string             `bson:"created_at"`
		EditedAt  string             `bson:"edited_at"`
		Tags      []string           `bson:"tags"`
	}
	var docs []Document
	var notes []note.Note
	ctx := context.Background()
	collection := client.Database("note").Collection("notes")
	log.Info("start: ", time.Now())
	filter := bson.M{"content": primitive.Regex{Pattern: kw, Options: ""}}
	if tag != "" {
		filter["tags"] = bson.M{"$in": []string{tag}}
	}
	cursor, err := collection.Find(
		ctx,
		filter,
		options.SetSort(bson.M{"edited_at": -1}),
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
		created, _ := time.Parse(time.RFC3339, d.CreatedAt)
		edited, _ := time.Parse(time.RFC3339, d.CreatedAt)
		notes = append(notes, note.Note{
			ID:        d.ID.Hex(),
			Content:   d.Content,
			Comment:   d.Comment,
			CreatedAt: created,
			EditedAt:  edited,
			Tags:      d.Tags,
		})
	}
	return notes, nil
}

func (u *noteRepository) AddNotes(notes []note.Note) ([]string, error) {
	var res []string
	type Document struct {
		Content   string   `bson:"content"`
		Comment   string   `bson:"comment"`
		CreatedAt string   `bson:"created_at"`
		EditedAt  string   `bson:"edited_at"`
		Tags      []string `bson:"tags"`
	}
	var ds []interface{}
	for _, n := range notes {
		ds = append(ds, Document{
			Content:   n.Content,
			Comment:   n.Comment,
			CreatedAt: time.Now().Format(time.RFC3339),
			EditedAt:  time.Now().Format(time.RFC3339),
			Tags:      n.Tags,
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
		ID       primitive.ObjectID `bson:"_id"`
		Content  string             `bson:"content"`
		Comment  string             `bson:"comment"`
		EditedAt string             `bson:"edited_at"`
		Tags     []string           `bson:"tags"`
	}
	ctx := context.Background()
	idPrimitive, err := primitive.ObjectIDFromHex(n.ID)
	d := Document{
		ID:       idPrimitive,
		Content:  n.Content,
		Comment:  n.Comment,
		EditedAt: time.Now().Format(time.RFC3339),
		Tags:     n.Tags,
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

func (u *noteRepository) GetTags() ([]entity.Tag, error) {
	type Doc struct {
		Name  string `bson:"_id"`
		Count int
	}
	var docs []Doc
	ctx := context.Background()
	collection := client.Database("note").Collection("notes")
	unwindStage := bson.D{
		{Key: "$unwind", Value: "$tags"},
	}
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$tags"},
			{Key: "count", Value: bson.D{
				{Key: "$sum", Value: 1},
			}},
		}},
	}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{unwindStage, groupStage})
	if err != nil {
		log.Error(err)
		return []entity.Tag{}, err
	}
	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	var tags []entity.Tag
	for _, d := range docs {
		tags = append(tags, entity.Tag{
			Name:  d.Name,
			Count: d.Count,
		})
	}
	return tags, nil
}
