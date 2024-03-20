package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timeoutDuration = 5 * time.Second
	dbName          = "mundo"
	collName        = "wordbooks"
)

type impl struct {
	rw *mongo.Client
}
