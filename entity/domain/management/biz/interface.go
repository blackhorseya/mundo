//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	idA "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
)

// IManagementBiz is the interface for management biz.
type IManagementBiz interface {
	// CreateWordBook is used to create a word book.
	CreateWordBook(ctx contextx.Contextx, by *idA.Member, name string) (item *agg.Wordbook, err error)

	// GetWordBookByName is used to get word book by name.
	GetWordBookByName(ctx contextx.Contextx, name string) (item *agg.Wordbook, err error)
}
