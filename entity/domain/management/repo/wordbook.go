//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
)

// ListOption is the option for list wordbooks.
type ListOption struct {
	Limit  int
	Offset int
}

// IWordbookRepo is the interface for wordbook repository.
type IWordbookRepo interface {
	// Create is used to create a wordbook.
	Create(ctx contextx.Contextx, book *agg.Wordbook) (err error)

	// GetByName is used to get a wordbook by name.
	GetByName(ctx contextx.Contextx, name string) (item *agg.Wordbook, err error)

	// List is used to list wordbooks.
	List(ctx contextx.Contextx, opts ListOption) (items []*agg.Wordbook, total int, err error)

	// AddWord is used to add word to wordbook.
	AddWord(ctx contextx.Contextx, book *agg.Wordbook, word string) (err error)
}
