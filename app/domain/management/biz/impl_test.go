package biz

import (
	"testing"

	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl      *gomock.Controller
	wordbooks *repo.MockIWordbookRepo
	biz       biz.IManagementBiz
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.wordbooks = repo.NewMockIWordbookRepo(s.ctrl)
	s.biz = NewManagementBiz(s.wordbooks)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
