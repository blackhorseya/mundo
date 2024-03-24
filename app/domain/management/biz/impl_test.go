package biz

import (
	"errors"
	"reflect"
	"testing"

	idA "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/identity/model"
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	model2 "github.com/blackhorseya/mundo/entity/domain/management/model"
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
	member1 := &idA.Member{
		User: model.User{
			ID: "tester1",
			Profile: &model.Profile{
				Name:  "tester1",
				Email: "",
			},
		},
	}
	book1 := &agg.Wordbook{
		Collection: model2.Collection{
			ID:      "",
			Name:    "test_book1",
			OwnerID: member1.ID,
		},
	}

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
		{
			name: "create by name then error",
			args: args{by: member1, name: book1.Name, mock: func() {
				s.wordbooks.EXPECT().Create(gomock.Any(), book1).Return(errors.New("mock error")).Times(1)
			}},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "create by name then ok",
			args: args{by: member1, name: book1.Name, mock: func() {
				s.wordbooks.EXPECT().Create(gomock.Any(), book1).Return(nil).Times(1)
			}},
			wantItem: book1,
			wantErr:  false,
		},
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
	book1 := &agg.Wordbook{
		Collection: model2.Collection{
			ID:      "test_book1",
			Name:    "test_book1",
			OwnerID: "tester1",
		},
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
			name: "get by name then error",
			args: args{name: book1.Name, mock: func() {
				s.wordbooks.EXPECT().GetByName(gomock.Any(), book1.Name).Return(nil, errors.New("mock error")).Times(1)
			}},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "get by name then ok",
			args: args{name: book1.Name, mock: func() {
				s.wordbooks.EXPECT().GetByName(gomock.Any(), book1.Name).Return(book1, nil).Times(1)
			}},
			wantItem: book1,
			wantErr:  false,
		},
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
