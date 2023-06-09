package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/nokaka/Xray-core/app/dispatcher"
	_ "github.com/nokaka/Xray-core/app/proxyman/inbound"
	_ "github.com/nokaka/Xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/nokaka/Xray-core/app/commander"
	_ "github.com/nokaka/Xray-core/app/log/command"
	_ "github.com/nokaka/Xray-core/app/proxyman/command"
	_ "github.com/nokaka/Xray-core/app/stats/command"

	// Developer preview services
	_ "github.com/nokaka/Xray-core/app/observatory/command"

	// Other optional features.
	_ "github.com/nokaka/Xray-core/app/dns"
	_ "github.com/nokaka/Xray-core/app/dns/fakedns"
	_ "github.com/nokaka/Xray-core/app/log"
	_ "github.com/nokaka/Xray-core/app/metrics"
	_ "github.com/nokaka/Xray-core/app/policy"
	_ "github.com/nokaka/Xray-core/app/reverse"
	_ "github.com/nokaka/Xray-core/app/router"
	_ "github.com/nokaka/Xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/nokaka/Xray-core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/nokaka/Xray-core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/nokaka/Xray-core/proxy/blackhole"
	_ "github.com/nokaka/Xray-core/proxy/dns"
	_ "github.com/nokaka/Xray-core/proxy/dokodemo"
	_ "github.com/nokaka/Xray-core/proxy/freedom"
	_ "github.com/nokaka/Xray-core/proxy/http"
	_ "github.com/nokaka/Xray-core/proxy/loopback"
	_ "github.com/nokaka/Xray-core/proxy/mtproto"
	_ "github.com/nokaka/Xray-core/proxy/shadowsocks"
	_ "github.com/nokaka/Xray-core/proxy/socks"
	_ "github.com/nokaka/Xray-core/proxy/trojan"
	_ "github.com/nokaka/Xray-core/proxy/vless/inbound"
	_ "github.com/nokaka/Xray-core/proxy/vless/outbound"
	_ "github.com/nokaka/Xray-core/proxy/vmess/inbound"
	_ "github.com/nokaka/Xray-core/proxy/vmess/outbound"
	_ "github.com/nokaka/Xray-core/proxy/wireguard"

	// Transports
	_ "github.com/nokaka/Xray-core/transport/internet/domainsocket"
	_ "github.com/nokaka/Xray-core/transport/internet/grpc"
	_ "github.com/nokaka/Xray-core/transport/internet/http"
	_ "github.com/nokaka/Xray-core/transport/internet/kcp"
	_ "github.com/nokaka/Xray-core/transport/internet/quic"
	_ "github.com/nokaka/Xray-core/transport/internet/reality"
	_ "github.com/nokaka/Xray-core/transport/internet/tcp"
	_ "github.com/nokaka/Xray-core/transport/internet/tls"
	_ "github.com/nokaka/Xray-core/transport/internet/udp"
	_ "github.com/nokaka/Xray-core/transport/internet/websocket"

	// Transport headers
	_ "github.com/nokaka/Xray-core/transport/internet/headers/http"
	_ "github.com/nokaka/Xray-core/transport/internet/headers/noop"
	_ "github.com/nokaka/Xray-core/transport/internet/headers/srtp"
	_ "github.com/nokaka/Xray-core/transport/internet/headers/tls"
	_ "github.com/nokaka/Xray-core/transport/internet/headers/utp"
	_ "github.com/nokaka/Xray-core/transport/internet/headers/wechat"
	_ "github.com/nokaka/Xray-core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/nokaka/Xray-core/main/json"
	_ "github.com/nokaka/Xray-core/main/toml"
	_ "github.com/nokaka/Xray-core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/nokaka/Xray-core/main/confloader/external"

	// Commands
	_ "github.com/nokaka/Xray-core/main/commands/all"
)
