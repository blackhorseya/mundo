package cmds

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// GetBookByNameCommand is a command to get book by name.
type GetBookByNameCommand struct {
	mgmt biz.IManagementBiz
}

func (cmd *GetBookByNameCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	// todo: 2024/3/25|sean|implement get book by name command
	return nil, nil
}
