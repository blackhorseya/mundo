package agg

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/blackhorseya/mundo/entity/domain/management/model"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

//go:embed book.tmpl
var f embed.FS

// Wordbook is an aggregate for word book.
type Wordbook struct {
	model.Collection
}

// FlexContainer will return the flex container message.
func (x *Wordbook) FlexContainer() (messaging_api.MessageInterface, error) {
	tmpl, err := template.New("book.tmpl").ParseFS(f, "book.tmpl")
	if err != nil {
		return nil, err
	}

	var layout bytes.Buffer
	err = tmpl.Execute(&layout, x)
	if err != nil {
		return nil, err
	}

	return messaging_api.UnmarshalFlexContainer(layout.Bytes())
}
