package cmds

import (
	"fmt"
	"strings"

	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type HelpCommand struct {
}

func (cmd *HelpCommand) Help(ctx contextx.Contextx) string {
	return fmt.Sprintf("Usage: %s\tshow help", "/help")
}

func (cmd *HelpCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text != "/help" {
		return nil, nil
	}

	commands := NewCommands(nil, nil)
	var helps []string
	for _, command := range commands {
		helps = append(helps, command.Help(ctx))
	}

	return []messaging_api.MessageInterface{
		&messaging_api.TextMessage{
			Text: strings.Join(helps, "\n"),
		},
	}, nil
}
