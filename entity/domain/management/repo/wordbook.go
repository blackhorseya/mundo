//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/mundo/entity/domain/management/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
)

// IWordbookRepo is the interface for wordbook repository.
type IWordbookRepo interface {
	// Create is used to create a wordbook.
	Create(ctx contextx.Contextx, book *agg.Wordbook) (err error)

	// GetByID is used to get a wordbook by id.
	GetByID(ctx contextx.Contextx, id string) (item *agg.Wordbook, err error)
}
