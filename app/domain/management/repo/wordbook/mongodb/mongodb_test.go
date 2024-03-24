package mongodb

import (
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/model"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/blackhorseya/mundo/pkg/storage/mongodbx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteTester struct {
	suite.Suite

	container *mongodbx.Container
	rw        *mongo.Client
	repo      repo.IWordbookRepo
}

func (s *suiteTester) SetupTest() {
	ctx := contextx.Background()

	container, err := mongodbx.NewContainer(ctx)
	s.Require().NoError(err)
	s.container = container

	dsn, err := container.ConnectionString(ctx)
	s.Require().NoError(err)

	rw, err := mongodbx.NewClientWithDSN(dsn)
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewWordbookRepo(rw)
}

func (s *suiteTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.container != nil {
		_ = s.container.Terminate(contextx.Background())
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_Create() {
	book1 := &agg.Wordbook{
		Collection: model.Collection{
			ID:      "",
			Name:    "book1",
			OwnerID: "tester1",
		},
	}

	type args struct {
		ctx  contextx.Contextx
		book *agg.Wordbook
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "create then ok",
			args:    args{book: book1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Create(tt.args.ctx, tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *suiteTester) Test_impl_GetByName() {
	book1 := &wordbook{
		ID:      primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:    "test_book1",
		OwnerID: "tester1",
	}

	type args struct {
		ctx  contextx.Contextx
		name string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Wordbook
		wantErr  bool
	}{
		{
			name: "ok",
			args: args{name: book1.Name, mock: func() {
				_, _ = s.rw.Database(dbName).Collection(collName).InsertOne(contextx.Background(), book1)
			}},
			wantItem: book1.ToAgg(),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.repo.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("GetByName() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
