package cmds

import (
	"strings"

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
	if strings.HasPrefix(text, "/create ") {
		name := strings.TrimPrefix(text, "/create ")

		// todo: 2024/3/21|sean|implement me
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: name,
			},
		}, nil
	}

	return nil, nil
}
