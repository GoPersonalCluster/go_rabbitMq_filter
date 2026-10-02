package vo

import (
	"fmt"
	"net"
)

type IPAddress struct {
	value string
}

func NewIPAddress(value string) (IPAddress, error) {

	return IPAddress{
		value: value,
	}, nil
}
func (ip IPAddress) IsValid(value string) error {
	if value == "" {
		return fmt.Errorf("IP address cannot be empty")
	}

	ip := net.ParseIP(value)

}
func (ip IPAddress) Value() string {
	return ip.value
}

func (ip IPAddress) String() string {
	return ip.value
}
