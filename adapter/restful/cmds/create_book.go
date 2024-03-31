package cmds

import (
	"fmt"
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

func (cmd *CreateBookCommand) Help(ctx contextx.Contextx) string {
	return fmt.Sprintf("Usage: %s\tcreate a book\n", "/create {name}")
}

func (cmd *CreateBookCommand) Execute(
	ctx contextx.Contextx,
	who *agg.Member,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "/create ") {
		name := strings.TrimPrefix(text, "/create ")

		got, err := cmd.mgmt.CreateWordBook(ctx, who, name)
		if err != nil {
			return nil, err
		}

		container, err := got.FlexContainer()
		if err != nil {
			return nil, err
		}

		return []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  got.Name,
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}
