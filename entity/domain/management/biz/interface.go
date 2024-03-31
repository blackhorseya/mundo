//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	idA "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
)

// ListWordBooksOption is the option for list word books.
type ListWordBooksOption struct {
	Page int
	Size int
}

// IManagementBiz is the interface for management biz.
type IManagementBiz interface {
	// GetWordBookByName is used to get word book by name.
	GetWordBookByName(ctx contextx.Contextx, name string) (item *agg.Wordbook, err error)

	// ListWordBooks is used to list word books.
	ListWordBooks(
		ctx contextx.Contextx,
		by *idA.Member,
		opt ListWordBooksOption,
	) (items []*agg.Wordbook, total int, err error)

	// CreateWordBook is used to create a word book.
	CreateWordBook(ctx contextx.Contextx, by *idA.Member, name string) (item *agg.Wordbook, err error)

	// AddWordToBook is used to add word to book.
	AddWordToBook(ctx contextx.Contextx, by *idA.Member, bookName string, word string) (item *agg.Wordbook, err error)
}
