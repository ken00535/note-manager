package db

import (
	"context"
	"fmt"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/logger"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	log    logger.Logger
)

// Init connect to DB with parameter written in config file
func Init(logInst logger.Logger) {
	log = logInst
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Print(evt.Command)
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := fmt.Sprintf("mongodb://%v:%v", config.GetDbAddress(), config.GetDbPort())
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url).SetMonitor(cmdMonitor))
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
