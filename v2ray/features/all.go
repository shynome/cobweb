package features

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/v2fly/v2ray-core/v4/app/dispatcher"
	_ "github.com/v2fly/v2ray-core/v4/app/proxyman/inbound"
	_ "github.com/v2fly/v2ray-core/v4/app/proxyman/outbound"

	// Other optional features.
	_ "github.com/v2fly/v2ray-core/v4/app/dns"
	_ "github.com/v2fly/v2ray-core/v4/app/dns/fakedns"
	_ "github.com/v2fly/v2ray-core/v4/app/log"
	_ "github.com/v2fly/v2ray-core/v4/app/policy"
	_ "github.com/v2fly/v2ray-core/v4/app/router"
	_ "github.com/v2fly/v2ray-core/v4/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/tagged/taggedimpl"

	// Inbound and outbound proxies.
	_ "github.com/v2fly/v2ray-core/v4/proxy/blackhole"
	_ "github.com/v2fly/v2ray-core/v4/proxy/dns"
	_ "github.com/v2fly/v2ray-core/v4/proxy/freedom"
	_ "github.com/v2fly/v2ray-core/v4/proxy/vmess/inbound"

	// Transports
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/websocket"

	// Geo loaders
	_ "github.com/v2fly/v2ray-core/v4/infra/conf/geodata/memconservative"
	_ "github.com/v2fly/v2ray-core/v4/infra/conf/geodata/standard"
)
