package cmds

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// AddWordToBookCommand is a command to add a word to a book.
type AddWordToBookCommand struct {
	mgmt biz.IManagementBiz
}

func (cmd *AddWordToBookCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	// todo: 2024/3/25|sean|implement add word to book command
	return nil, nil
}
