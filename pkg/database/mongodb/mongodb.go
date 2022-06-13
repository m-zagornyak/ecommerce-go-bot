package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {

}
