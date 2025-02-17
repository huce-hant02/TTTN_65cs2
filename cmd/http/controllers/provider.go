package controllers

import (
	"github.com/google/wire"
	edit_history "mono-base/cmd/http/controllers/edit-history"
	"mono-base/cmd/http/controllers/user"
)

var ControllerProviders = wire.NewSet(
	user.NewUserControllerV1,
	user.NewUserControllerV2,
	edit_history.NewEditHistoryControllerV1,
)
