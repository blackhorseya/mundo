package cmds

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// CreateBookCommand is the command for create book.
type CreateBookCommand struct {
	mgmt biz.IManagementBiz
}

func (cmd *CreateBookCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	// todo: 2024/3/21|sean|implement me
	return nil, nil
}
