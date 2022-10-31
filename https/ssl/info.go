package ssl

import (
	"crypto/tls"
	"crypto/x509/pkix"
	"errors"
	"net"
	"time"
)

// DomainSslInfo 域名SSL证书信息
type DomainSslInfo struct {
	Name       string    `json:"name"`
	Issuer     pkix.Name `json:"issuer"`
	IsCA       bool      `json:"is_ca"`
	CreateTime time.Time `json:"create_time"`
	ExpiryTime time.Time `json:"expiry_time"`
}

// GetDomainSslInformation 获取域名证书信息
func GetDomainSslInformation(domain string, port ...string) (data DomainSslInfo, err error) {
	p := "443"
	if len(port) > 0 {
		p = port[0]
	}
	conn, err := tls.DialWithDialer(&net.Dialer{
		Timeout:  time.Second * 5,
		Deadline: time.Now().Add(time.Second * 9),
	}, "tcp", domain+":"+p, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return
	}
	defer conn.Close()
	stats := conn.ConnectionState()
	certs := stats.PeerCertificates
	if len(certs) == 0 {
		err = errors.New("Domain ssl info get fail.")
		return
	}
	info := certs[0]
	data.Name = info.DNSNames[0]
	data.Issuer = info.Issuer
	data.IsCA = info.IsCA
	data.CreateTime = info.NotBefore
	data.ExpiryTime = info.NotAfter
	return
}
