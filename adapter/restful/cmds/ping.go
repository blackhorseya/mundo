package cmds

import (
	"fmt"

	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// PingCommand is the command for ping.
type PingCommand struct {
}

func (cmd *PingCommand) Help(ctx contextx.Contextx) string {
	return fmt.Sprintf("Usage: %s\tping command", "/ping")
}

func (cmd *PingCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "/ping" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "pong",
			},
		}, nil
	}

	return nil, nil
}
