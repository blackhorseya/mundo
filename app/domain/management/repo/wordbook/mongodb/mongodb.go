package mongodb

import (
	"time"

	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
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

// NewWordbookRepo is used to create a new wordbook repository.
func NewWordbookRepo(rw *mongo.Client) repo.IWordbookRepo {
	return &impl{rw: rw}
}

func (i *impl) Create(ctx contextx.Contextx, book *agg.Wordbook) (err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	created := newWordbook(book)
	coll := i.rw.Database(dbName).Collection(collName)
	_, err = coll.InsertOne(timeout, created)
	if err != nil {
		return err
	}

	book.ID = created.ID.Hex()
	return nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (item *agg.Wordbook, err error) {
	// todo: 2024/3/21|sean|implement me
	panic("implement me")
}
