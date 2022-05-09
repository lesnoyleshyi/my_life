package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongoConnection(connstr string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(connstr)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, fmt.Errorf("unable to connect mongo: %w", err)
	}
	return client, nil
}
