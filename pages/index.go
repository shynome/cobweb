package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetDashBoard(ctx *context.Context) (types.Panel, error) {
	return types.Panel{
		Content:     "hello world",
		Title:       "Dashboard",
		Description: "dashboard example",
	}, nil
}
