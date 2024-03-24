package agg

import (
	"github.com/blackhorseya/mundo/entity/domain/management/model"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// Wordbook is an aggregate for word book.
type Wordbook struct {
	model.Collection
}

// FlexContainer will return the flex container message.
func (x *Wordbook) FlexContainer() ([]messaging_api.MessageInterface, error) {
	// todo: 2024/3/25|sean|implement me
	return nil, nil
}
