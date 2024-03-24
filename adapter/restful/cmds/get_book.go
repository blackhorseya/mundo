package cmds

import (
	"errors"
	"strings"

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
	if !strings.HasPrefix(text, "/get_book ") {
		return nil, nil
	}

	name := strings.TrimSpace(strings.TrimPrefix(text, "/get_book "))
	if name == "" {
		return nil, errors.New("name is required")
	}

	got, err := cmd.mgmt.GetWordBookByName(ctx, name)
	if err != nil {
		return nil, err
	}

	container, err := got.FlexContainer()
	if err != nil {
		return nil, err
	}

	return []messaging_api.MessageInterface{
		&messaging_api.FlexMessage{
			AltText:  "Book " + got.Name,
			Contents: container,
		},
	}, nil
}
