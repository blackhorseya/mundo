package cmds

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/entity/domain/management/biz"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/blackhorseya/mundo/pkg/openaix"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	// Execute is the function to execute the command.
	Execute(ctx contextx.Contextx, who *agg.Member, text string) ([]messaging_api.MessageInterface, error)

	// Help is the function to get the help message.
	Help(ctx contextx.Contextx) string
}

// NewCommands is the function to create the text commands.
func NewCommands(client *openaix.Client, mgmt biz.IManagementBiz) []TextCommander {
	return []TextCommander{
		&HelpCommand{},
		&PingCommand{},
		&WhoAmICommand{},
		&ExplainCommand{client: client},
		&CreateBookCommand{mgmt: mgmt},
		&GetBookByNameCommand{mgmt: mgmt},
		&AddWordToBookCommand{mgmt: mgmt},
	}
}
