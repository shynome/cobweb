package v2ray

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ContextKey string

const ContextV2rayKey ContextKey = "v2"

func FromContext(ctx context.Context) V2rayControl {
	return ctx.Value(ContextV2rayKey).(V2rayControl)
}

func InjectV2rayMiddleware(v2 *V2ray) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), ContextV2rayKey, v2)
			r := c.Request().WithContext(ctx)
			c.SetRequest(r)
			return next(c)
		}
	}
}
