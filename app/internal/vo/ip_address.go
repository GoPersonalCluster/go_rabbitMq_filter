package vo

import (
	"errors"
	"net"
)

type IPAddress struct {
	value net.IP
}

func NewIPAddress(value string) (IPAddress, error) {
	ip := net.ParseIP(value)

	if ip == nil {
		return IPAddress{}, errors.New("invalid IP address")
	}

	return IPAddress{
		value: ip,
	}, nil
}

func (ip IPAddress) String() string {
	return ip.value.String()
}
