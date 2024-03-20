package biz

import (
	idA "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/entity/domain/management/repo"
	"github.com/blackhorseya/mundo/pkg/contextx"
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

func (i *impl) CreateWordBook(ctx contextx.Contextx, by *idA.Member, name string) (item *agg.Wordbook, err error) {
	// todo: 2024/3/21|sean|implement me
	panic("implement me")
}

func (i *impl) GetWordBookByID(ctx contextx.Contextx, id string) (item *agg.Wordbook, err error) {
	// todo: 2024/3/21|sean|implement me
	panic("implement me")
}
