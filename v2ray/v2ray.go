package v2ray

import (
	"context"
	"fmt"

	"github.com/shynome/cobweb/models"
	"github.com/shynome/cobweb/v2ray/config"
	_ "github.com/shynome/cobweb/v2ray/features"
	core "github.com/v2fly/v2ray-core/v4"
	"github.com/v2fly/v2ray-core/v4/features/inbound"
	"github.com/v2fly/v2ray-core/v4/proxy"
)

type V2rayControl interface {
	AddUser(user models.V2rayUser) error
	RemoveUser(id string) error
}

type V2ray struct {
	um     proxy.UserManager
	Server *core.Instance
}

func New(users []models.V2rayUser) (v2 *V2ray, err error) {
	v2 = &V2ray{}
	cfg := config.GenConfig()

	server, err := core.New(cfg)
	if err != nil {
		return
	}
	v2.Server = server

	um, err := getUserManager(server)
	if err != nil {
		return
	}
	v2.um = um
	for _, u := range users {
		if err = v2.AddUser(u); err != nil {
			return
		}
	}

	return
}

func getUserManager(c *core.Instance) (um proxy.UserManager, err error) {
	im := c.GetFeature(inbound.ManagerType()).(inbound.Manager)
	var ih inbound.Handler
	if ih, err = im.GetHandler(context.Background(), "v2ws"); err != nil {
		return
	}
	gi, ok := ih.(proxy.GetInbound)
	if !ok {
		err = fmt.Errorf("can't get inbound proxy from Handler Manager")
		return
	}
	p := gi.GetInbound()
	um, ok = p.(proxy.UserManager)
	if !ok {
		err = fmt.Errorf("proxy is not a UserManager")
		return
	}
	return
}

func (v2 *V2ray) AddUser(user models.V2rayUser) (err error) {
	u, err := config.ToMemoryUser(user)
	if err != nil {
		return
	}
	return v2.um.AddUser(context.Background(), u)
}

func (v2 *V2ray) RemoveUser(id string) (err error) {
	return v2.um.RemoveUser(context.Background(), id)
}

var _ V2rayControl = &V2ray{}
