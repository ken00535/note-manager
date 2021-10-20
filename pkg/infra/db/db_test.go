package db

import (
	"context"
	"fmt"
	"testing"

	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/logger"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConnect(t *testing.T) {
	filter := bson.D{}
	type Note struct {
		Content string
		Comment string
	}
	var result []Note
	log := logger.NewMockLogger()
	config.Init(log)
	Init(log)
	collection := client.Database("note").Collection("notes")
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	fmt.Print(result)
	assert.NoError(t, err)
}

func TestConnectFindAll(t *testing.T) {
	filter := bson.D{}
	type Note struct {
		Content string
		Comment string
	}
	var result []Note
	log := logger.NewMockLogger()
	config.Init(log)
	Init(log)
	collection := client.Database("note").Collection("notes")
	ctx := context.Background()
	cursor, err := collection.Find(ctx, filter)
	cursor.All(ctx, &result)
	fmt.Print(result)
	assert.NoError(t, err)
}
