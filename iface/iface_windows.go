package iface

import (
	"fmt"

	"github.com/pion/transport/v3"

	"yiji.one/punch/iface/bind"
	"yiji.one/punch/iface/netstack"
)

// NewWGIFace Creates a new WireGuard interface instance
func NewWGIFace(iFaceName string, address string, wgPort int, wgPrivKey string, mtu int, transportNet transport.Net, args *MobileIFaceArguments, filterFn bind.FilterFn) (*WGIface, error) {
	wgAddress, err := parseWGAddress(address)
	if err != nil {
		return nil, err
	}

	wgIFace := &WGIface{
		userspaceBind: true,
	}

	if netstack.IsEnabled() {
		wgIFace.tun = newTunNetstackDevice(iFaceName, wgAddress, wgPort, wgPrivKey, mtu, transportNet, netstack.ListenAddr(), filterFn)
		return wgIFace, nil
	}

	wgIFace.tun = newTunDevice(iFaceName, wgAddress, wgPort, wgPrivKey, mtu, transportNet, filterFn)
	return wgIFace, nil
}

// CreateOnAndroid this function make sense on mobile only
func (w *WGIface) CreateOnAndroid([]string, string, []string) error {
	return fmt.Errorf("this function has not implemented on non mobile")
}

// GetInterfaceGUIDString returns an interface GUID. This is useful on Windows only
func (w *WGIface) GetInterfaceGUIDString() (string, error) {
	return w.tun.(*tunDevice).getInterfaceGUIDString()
}
