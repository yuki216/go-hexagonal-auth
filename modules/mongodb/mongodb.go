package mongodb

import (
	"context"
	"go-hexagonal-auth/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBServer struct {
	cfg *config.Config
}

type MongoDBServerInterface interface {
	Connect(ctx context.Context) (*mongo.Database, error)
}

func NewMongoDBServer(cfg *config.Config) MongoDBServerInterface {
	return &mongoDBServer{
		cfg: cfg,
	}
}

func (r *mongoDBServer) Connect(ctx context.Context) (*mongo.Database, error) {
	timeout := time.Duration(r.cfg.Server.WriteTimeout) * time.Second
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	clientOptions.ConnectTimeout = &timeout
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}
