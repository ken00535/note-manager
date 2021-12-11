package db

import (
	"context"
	"crypto/tls"
	"fmt"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	log    logger.Logger
	once   sync.Once
	c      config.Config
)

// Init connect to DB with parameter written in config file
func Init(logInst logger.Logger) {
	once.Do(func() {
		log = logInst
		c = config.Init()
	})
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Info(evt.Command)
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := c.GetDbOption()
	url := fmt.Sprintf("mongodb://%v:%v", opt.Address, opt.Port)
	auth := options.Credential{
		AuthMechanism: opt.Mechanism,
		Username:      opt.Username,
		Password:      opt.Password,
	}
	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	var err error
	clientOpt := options.Client().
		ApplyURI(url).
		SetMonitor(cmdMonitor).
		SetAuth(auth).
		SetTLSConfig(tlsCfg)
	client, err = mongo.Connect(ctx, clientOpt)
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
