package cmds

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// WhoAmICommand is the command for whoami.
type WhoAmICommand struct {
}

func (cmd *WhoAmICommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "whoami" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: who.ID,
			},
		}, nil
	}

	return nil, nil
}
