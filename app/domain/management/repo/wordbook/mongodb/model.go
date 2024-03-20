package mongodb

import (
	"time"

	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type wordbook struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	OwnerID string             `bson:"owner_id"`
}

func newWordbook(v *agg.Wordbook) *wordbook {
	return &wordbook{
		ID:      primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:    v.Name,
		OwnerID: v.OwnerID,
	}
}
