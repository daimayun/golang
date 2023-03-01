package function

import (
	"errors"
	"net"
)

// GetLocalIp 获取本机IP
func GetLocalIp() (net.IP, error) {
	is, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range is {
		if i.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if i.Flags&net.FlagLoopback != 0 {
			continue // loop back interface
		}
		var addresses []net.Addr
		addresses, err = i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addresses {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}
	return ip
}
