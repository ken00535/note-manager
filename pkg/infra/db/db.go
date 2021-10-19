package db

import (
	"context"
	"fmt"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/logger"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres" // import driver
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	log    logger.Logger
)

// Connect to DB with parameter written in config file
func Connect(logInst logger.Logger) {
	log = logInst
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := fmt.Sprintf("mongodb://%v:%v", config.GetDbAddress(), config.GetDbPort())
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Panic(err)
	}
}

// Close session
func Close() {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Panic(err)
	}
}

// NewClient new a db client
func NewClient() *mongo.Client {
	return client
}
