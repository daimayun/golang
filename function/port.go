package function

import "net"

// GetFreePort 获取随机端口
func GetFreePort() (port int, err error) {
	var addr *net.TCPAddr
	addr, err = net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return
	}
	var listen *net.TCPListener
	listen, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port, nil
}
