package db

import (
	"context"
	"testing"

	"note-manager/pkg/config"
	"note-manager/pkg/infra/logger"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConnect(t *testing.T) {
	filter := bson.D{}
	var result bson.D
	log := logger.NewMockLogger()
	config.Init(log)
	Connect(log)
	collection := client.Database("note").Collection("notes")
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	assert.NoError(t, err)
}
