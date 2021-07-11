package config

import (
	"strconv"
	"strings"

	core "github.com/v2fly/v2ray-core/v4"
	"github.com/v2fly/v2ray-core/v4/app/dispatcher"
	"github.com/v2fly/v2ray-core/v4/app/dns"
	"github.com/v2fly/v2ray-core/v4/app/log"
	"github.com/v2fly/v2ray-core/v4/app/proxyman"
	"github.com/v2fly/v2ray-core/v4/app/router"
	logLevel "github.com/v2fly/v2ray-core/v4/common/log"
	"github.com/v2fly/v2ray-core/v4/common/net"
	"github.com/v2fly/v2ray-core/v4/common/serial"
	"github.com/v2fly/v2ray-core/v4/proxy/blackhole"
	"github.com/v2fly/v2ray-core/v4/proxy/freedom"
)

var blackoutList = []string{
	"0.0.0.0/8",
	"10.0.0.0/8",
	"100.64.0.0/10",
	"127.0.0.0/8",
	"169.254.0.0/16",
	"172.16.0.0/12",
	"192.0.0.0/24",
	"192.0.2.0/24",
	"192.168.0.0/16",
	"198.18.0.0/15",
	"198.51.100.0/24",
	"203.0.113.0/24",
	"::1/128",
	"fc00::/7",
	"fe80::/10",
}

func initAppRouteBlackListRule() *router.RoutingRule {
	blackList := []*router.CIDR{}
	for _, filed := range blackoutList {

		match := strings.Split(filed, "/")
		ipStr, prefixStr := match[0], match[1]
		if ipStr == "" || prefixStr == "" {
			continue
		}
		ip := net.ParseAddress(ipStr).IP()
		prefix, err := strconv.ParseInt(prefixStr, 10, 32)
		if err != nil {
			continue
		}

		rule := &router.CIDR{
			Ip:     ip,
			Prefix: uint32(prefix),
		}
		blackList = append(blackList, rule)
	}
	return &router.RoutingRule{
		Geoip: []*router.GeoIP{
			{Cidr: blackList},
		},
		TargetTag: &router.RoutingRule_Tag{
			Tag: "blockout",
		},
	}
}

func initAppRoute() *serial.TypedMessage {

	rule := []*router.RoutingRule{}

	if len(blackoutList) != 0 {
		blackListRule := initAppRouteBlackListRule()
		rule = append(rule, blackListRule)
	}

	return serial.ToTypedMessage(&router.Config{Rule: rule})
}

func initAppDNS() *serial.TypedMessage {
	return serial.ToTypedMessage(&dns.Config{
		NameServer: []*dns.NameServer{
			{
				Address: &net.Endpoint{
					Address: &net.IPOrDomain{Address: &net.IPOrDomain_Ip{Ip: []byte{8, 8, 8, 8}}},
				},
			},
			{
				Address: &net.Endpoint{
					Address: &net.IPOrDomain{Address: &net.IPOrDomain_Ip{Ip: []byte{8, 8, 4, 4}}},
				},
			},
			{
				Address: &net.Endpoint{
					// Address: &net.IPOrDomain{Address: &net.IPOrDomain_Domain{Domain: "localhost"}},
					Address: &net.IPOrDomain{Address: &net.IPOrDomain_Ip{Ip: []byte{127, 0, 0, 1}}},
				},
			},
		},
	})
}

func initApp() []*serial.TypedMessage {
	// 设置路由
	routeService := initAppRoute()
	// dns
	dnsService := initAppDNS()
	// 设置日志
	logService := serial.ToTypedMessage(&log.Config{
		ErrorLogLevel: logLevel.Severity_Error,
		ErrorLogType:  log.LogType_Console,
	})
	return []*serial.TypedMessage{
		dnsService,
		routeService,
		logService,
		// init
		serial.ToTypedMessage(&dispatcher.Config{}),
		serial.ToTypedMessage(&proxyman.InboundConfig{}),
		serial.ToTypedMessage(&proxyman.OutboundConfig{}),
	}
}

func initInbound() []*core.InboundHandlerConfig {
	return []*core.InboundHandlerConfig{}
}

func initOutbound() []*core.OutboundHandlerConfig {
	return []*core.OutboundHandlerConfig{
		{
			Tag:           "direct",
			ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
		},
		{
			Tag: "blockout",
			ProxySettings: serial.ToTypedMessage(&blackhole.Config{
				Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
			}),
		},
	}
}

func getV2rayConfig() *core.Config {
	app := initApp()
	inbound := initInbound()
	outbound := initOutbound()
	config := &core.Config{
		App:      app,
		Inbound:  inbound,
		Outbound: outbound,
	}
	return config
}
