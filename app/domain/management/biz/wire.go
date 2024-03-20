package biz

import (
	"github.com/blackhorseya/mundo/app/domain/management/repo/wordbook/mongodb"
	"github.com/google/wire"
)

// ProvideManagementBiz is used to provide a new management biz.
var ProvideManagementBiz = wire.NewSet(
	NewManagementBiz,
	mongodb.NewWordbookRepo,
)
