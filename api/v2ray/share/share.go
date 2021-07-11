package share

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/shynome/cobweb/config"
	qrsvg "github.com/wamuir/svg-qr-code"
)

type Params struct {
	UUID string `form:"uuid"`
}

func Register(g *echo.Group) {

	g.POST("/link", func(c echo.Context) (err error) {
		var params Params
		if err = c.Bind(&params); err != nil {
			return
		}
		link, err := genShareLink(c, params.UUID)
		if err != nil {
			return
		}
		return c.String(http.StatusOK, link)
	})
	g.POST("/qrcode", func(c echo.Context) (err error) {
		var params Params
		if err = c.Bind(&params); err != nil {
			return
		}
		link, err := genShareLink(c, params.UUID)
		if err != nil {
			return
		}
		qr, err := qrsvg.New(link)
		qr.Borderwidth = 8
		qr.Blocksize = 6
		if err != nil {
			return
		}
		return c.Blob(http.StatusOK, "image/svg+xml", []byte(qr.String()))
	})
	g.POST("/vnext", func(c echo.Context) (err error) {
		var params Params
		if err = c.Bind(&params); err != nil {
			return
		}
		vnext, err := genShareVNext(c, params.UUID)
		if err != nil {
			return
		}
		return c.JSONPretty(http.StatusOK, vnext, "  ")
	})
}

func genShareVNext(c echo.Context, uuid string) (VNEXT, error) {
	cfg := config.GetV2rayConfig()
	req := c.Request()
	requ := strings.Split(req.Host, ":")

	domain := cfg.UseDomain
	if domain == "" {
		domain = requ[0]
	}
	port := cfg.UsePort
	if port == "" {
		if len(requ) == 2 {
			port = requ[1]
		} else {
			port = "80"
		}
	}
	path := cfg.UsePath
	tls := cfg.UseTLS
	vnext := VNEXT{
		Version: "2",
		Remark:  "",
		Address: domain,
		Port:    port,
		ID:      uuid,
		AlertID: 0,
		Network: "ws",
		Type:    "",
		Host:    "",
		Path:    path,
		TLS:     tls,
	}
	return vnext, nil
}

func genShareLink(c echo.Context, uuid string) (string, error) {
	vnext, err := genShareVNext(c, uuid)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(vnext)
	if err != nil {
		return "", err
	}
	link := base64.StdEncoding.EncodeToString(b)
	return link, nil
}
