package v2ray

import (
	"github.com/labstack/echo/v4"
	"github.com/shynome/cobweb/api/v2ray/share"
)

func Register(g *echo.Group) {
	share.Register(g.Group("/share"))
}
