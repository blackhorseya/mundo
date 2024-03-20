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
	CreateWordBook(ctx contextx.Contextx, by *idA.Member, name string) (item *agg.WordBook, err error)

	// GetWordBookByID is used to get a word book by id.
	GetWordBookByID(ctx contextx.Contextx, id string) (item *agg.WordBook, err error)
}
