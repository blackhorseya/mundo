package mongodb

import (
	"testing"

	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/blackhorseya/mundo/pkg/storage/mongodbx"
	"github.com/stretchr/testify/suite"
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
