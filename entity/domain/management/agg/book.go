package agg

import (
	"github.com/blackhorseya/mundo/entity/domain/management/model"
)

// WordBook is an aggregate for word book.
type WordBook struct {
	model.Collection
}
