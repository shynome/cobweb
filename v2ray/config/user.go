package config

import (
	"github.com/shynome/cobweb/models"
	"github.com/v2fly/v2ray-core/v4/common/protocol"
	"github.com/v2fly/v2ray-core/v4/common/serial"
	"github.com/v2fly/v2ray-core/v4/proxy/vmess"
)

func ToMemoryUser(u models.V2rayUser) (user *protocol.MemoryUser, err error) {
	va := vmess.Account{
		Id:      u.Uuid,
		AlterId: 1,
	}
	ma, err := va.AsAccount()
	if err != nil {
		return
	}
	user = &protocol.MemoryUser{
		Email:   u.Username,
		Level:   0,
		Account: ma,
	}
	return
}
func ToUser(u models.V2rayUser) (user *protocol.User, err error) {
	va := vmess.Account{
		Id:      u.Uuid,
		AlterId: 1,
	}
	user = &protocol.User{
		Email:   u.Username,
		Level:   0,
		Account: serial.ToTypedMessage(&va),
	}
	return
}
