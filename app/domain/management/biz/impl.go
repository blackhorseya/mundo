package biz

import (
	idA "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/entity/domain/management/model"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	wordbooks repo.IWordbookRepo
}

// NewManagementBiz is used to create a new management biz.
func NewManagementBiz(wordbooks repo.IWordbookRepo) biz.IManagementBiz {
	return &impl{
		wordbooks: wordbooks,
	}
}

func (i *impl) GetWordBookByName(ctx contextx.Contextx, name string) (item *agg.Wordbook, err error) {
	got, err := i.wordbooks.GetByName(ctx, name)
	if err != nil {
		ctx.Error("get a wordbook by name from database failed", zap.Error(err))
		return nil, err
	}

	return got, nil
}

func (i *impl) ListWordBooks(
	ctx contextx.Contextx,
	by *idA.Member,
	opt biz.ListWordBooksOption,
) (items []*agg.Wordbook, total int, err error) {
	// todo: 2024/3/31|sean|implement list word books
	panic("implement me")
}

func (i *impl) CreateWordBook(ctx contextx.Contextx, by *idA.Member, name string) (item *agg.Wordbook, err error) {
	ret := &agg.Wordbook{
		Collection: model.Collection{
			ID:      "",
			Name:    name,
			OwnerID: by.ID,
		},
	}
	err = i.wordbooks.Create(ctx, ret)
	if err != nil {
		ctx.Error("create a wordbook into database failed", zap.Error(err))
		return nil, err
	}

	return ret, nil
}

func (i *impl) AddWordToBook(
	ctx contextx.Contextx,
	by *idA.Member,
	bookName string,
	word string,
) (item *agg.Wordbook, err error) {
	// todo: 2024/3/31|sean|implement add word to book
	panic("implement me")
}
