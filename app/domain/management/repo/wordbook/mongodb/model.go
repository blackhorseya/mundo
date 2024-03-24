package mongodb

import (
	"time"

	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/model"
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

// ToAgg is used to convert wordbook to agg wordbook.
func (x *wordbook) ToAgg() *agg.Wordbook {
	return &agg.Wordbook{
		Collection: model.Collection{
			ID:      x.ID.Hex(),
			Name:    x.Name,
			OwnerID: x.OwnerID,
		},
	}
}
