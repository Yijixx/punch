package main

import (
	"context"
	"flag"

	log "github.com/sirupsen/logrus"
	"yiji.one/punch/client/internal"
)

func main() {
	flag.Parse()
	ctx := context.WithValue(context.TODO(), "k1", "v1")
	var ifaceName string
	var port int = 51822

	ic := internal.ConfigInput{
		ConfigPath: "",
		// NATExternalIPs NAT外部IP
		NATExternalIPs: make([]string, 0),
		// CustomDNSAddress 自定义DNS地址
		CustomDNSAddress: make([]byte, 0),
		// InterfaceName 接口名称
		InterfaceName: &ifaceName,
		// WireguardPort Wireguard端口
		WireguardPort: &port,
		// ExtraIFaceBlackList 额外接口黑名单
		ExtraIFaceBlackList: make([]string, 0),
	}

	config, err := internal.CreateInMemoryConfig(ic)
	if err != nil {
		log.Error("Error creating config: ", err)
		return
	}
	connectClient := internal.NewConnectClient(ctx, config)
	connectClient.Run()
}
