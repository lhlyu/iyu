package test

import (
	"fmt"
	"github.com/lhlyu/yutil"
	"net"
	"testing"
)

type A struct {
	List []string `json:"list"`
}

func TestSome(t *testing.T) {
	var a = "sda"
	var b int
	yutil.NotIgnore()
	yutil.JsonStrToObj(a, &b)

}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}
