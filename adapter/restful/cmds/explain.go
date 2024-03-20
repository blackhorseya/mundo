package cmds

import (
	"errors"
	"strings"

	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/query/model"
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
		prompt, err := model.GetPromptByString(model.ExplainTemplate, map[string]any{
			"word": word,
		})
		if err != nil {
			return nil, err
		}

		resp, err := cmd.client.CreateFunctionCall(ctx, prompt)
		if err != nil {
			return nil, err
		}

		if len(resp.Choices) == 0 {
			return nil, errors.New("no response")
		}

		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: resp.Choices[0].Message.Content,
			},
		}, nil
	}

	return nil, nil
}
