package config

import (
	core "github.com/v2fly/v2ray-core/v4"
	"github.com/v2fly/v2ray-core/v4/app/proxyman"
	"github.com/v2fly/v2ray-core/v4/common/net"
	"github.com/v2fly/v2ray-core/v4/common/protocol"
	"github.com/v2fly/v2ray-core/v4/common/serial"
	"github.com/v2fly/v2ray-core/v4/proxy/vmess/inbound"
	"github.com/v2fly/v2ray-core/v4/transport/internet"
	"github.com/shynome/cobweb/v2ray/websocket"
)

// GenConfig expose v2ray config
func GenConfig() *core.Config {

	wsPath := "/ray"
	wsPort := 3005

	usersInbound := &core.InboundHandlerConfig{
		Tag: "v2ws",
		ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
			PortRange: net.SinglePortRange(net.Port(wsPort)),
			StreamSettings: &internet.StreamConfig{
				Protocol: internet.TransportProtocol_WebSocket,
				TransportSettings: []*internet.TransportConfig{
					{
						Protocol: internet.TransportProtocol_WebSocket,
						Settings: serial.ToTypedMessage(&websocket.Config{
							Path: wsPath,
						}),
					},
				},
			},
		}),
		ProxySettings: serial.ToTypedMessage(&inbound.Config{
			User: []*protocol.User{},
		}),
	}

	config := getV2rayConfig()
	config.Inbound = append(config.Inbound, usersInbound)

	return config

}
