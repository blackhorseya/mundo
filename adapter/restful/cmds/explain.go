package cmds

import (
	"strings"

	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/blackhorseya/mundo/pkg/openaix"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// ExplainCommand is the command for explain.
type ExplainCommand struct {
	client *openaix.Client
}

func (cmd *ExplainCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "explain.") {
		word := strings.TrimPrefix(text, "explain.")
		// todo: 2024/3/20|sean|implement explain command

		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: word,
			},
		}, nil
	}

	return nil, nil
}
