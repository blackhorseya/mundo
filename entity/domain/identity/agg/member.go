package agg

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/model"
)

// Member is an aggregate root for member.
type Member struct {
	*model.User `json:",inline"`
}
