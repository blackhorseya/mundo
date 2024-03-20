package cmds

import (
	"github.com/blackhorseya/mundo/entity/domain/identity/agg"
	"github.com/blackhorseya/mundo/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	// Execute is the function to execute the command.
	Execute(ctx contextx.Contextx, who *agg.Member, text string) ([]messaging_api.MessageInterface, error)
}

// NewCommands is the function to create the text commands.
func NewCommands() []TextCommander {
	return []TextCommander{
		&PingCommand{},
		&WhoAmICommand{},
	}
}
