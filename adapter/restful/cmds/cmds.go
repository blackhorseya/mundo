package cmds

// TextCommander is the interface for text command.
type TextCommander interface {
}

// NewCommands is the function to create the text commands.
func NewCommands() []TextCommander {
	return []TextCommander{}
}
