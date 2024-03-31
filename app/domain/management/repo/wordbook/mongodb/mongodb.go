package mongodb

import (
	"time"

	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
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

func (i *impl) GetByName(ctx contextx.Contextx, name string) (item *agg.Wordbook, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	filter := bson.M{"name": name}
	coll := i.rw.Database(dbName).Collection(collName)
	var got *wordbook
	err = coll.FindOne(timeout, filter).Decode(&got)
	if err != nil {
		return nil, err
	}

	return got.ToAgg(), nil
}

func (i *impl) List(ctx contextx.Contextx, opts repo.ListOption) (items []*agg.Wordbook, total int, err error) {
	// todo: 2024/3/31|sean|implement list wordbooks
	panic("implement me")
}

func (i *impl) AddWord(ctx contextx.Contextx, book *agg.Wordbook, word string) (err error) {
	// todo: 2024/3/31|sean|implement add word to wordbook
	panic("implement me")
}
