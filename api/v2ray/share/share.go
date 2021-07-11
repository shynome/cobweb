package share

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/shynome/cobweb/config"
	qrsvg "github.com/wamuir/svg-qr-code"
)

type Params struct {
	UUID   string `form:"uuid"`
	Remark string `form:"remark"`
}

func Register(g *echo.Group) {

	g.POST("/link", func(c echo.Context) (err error) {
		var params Params
		if err = c.Bind(&params); err != nil {
			return
		}
		link, err := genShareLink(c, params)
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
		link, err := genShareLink(c, params)
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
		vnext, err := genShareVNext(c, params)
		if err != nil {
			return
		}
		return c.JSONPretty(http.StatusOK, vnext, "  ")
	})
}

func genShareVNext(c echo.Context, p Params) (VNEXT, error) {
	uuid := p.UUID
	remark := p.Remark

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

	var remarkPrefix string
	if cfg.UseRemarkPrefix != "" {
		remarkPrefix = cfg.UseRemarkPrefix
	} else {
		remarkPrefix = "ws"
		tls = ""
		if cfg.UseTLS == "tls" {
			tls = "s"
		}
		remarkPrefix = fmt.Sprintf("ws%s://%s:%s", tls, domain, port)
	}
	vnext := VNEXT{
		Version: "2",
		Remark:  remarkPrefix + " - " + remark,
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

func genShareLink(c echo.Context, p Params) (string, error) {
	vnext, err := genShareVNext(c, p)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(vnext)
	if err != nil {
		return "", err
	}
	link := base64.StdEncoding.EncodeToString(b)
	link = "vmess://" + link
	return link, nil
}
