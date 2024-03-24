package biz

import (
	"reflect"
	"testing"

	idA "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
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

func (s *suiteTester) Test_impl_CreateWordBook() {
	type args struct {
		ctx  contextx.Contextx
		by   *idA.Member
		name string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Wordbook
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.biz.CreateWordBook(tt.args.ctx, tt.args.by, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWordBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("CreateWordBook() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func (s *suiteTester) Test_impl_GetWordBookByName() {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.biz.GetWordBookByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWordBookByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("GetWordBookByName() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
