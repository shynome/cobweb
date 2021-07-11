package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shynome/cobweb/api/v2ray"
)

func Register(g *echo.Group) {
	v2ray.Register(g.Group("/v2ray"))
}
