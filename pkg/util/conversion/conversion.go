package conversion

import (
	"encoding/binary"
	"fmt"
	"net"
)

func IpToBinary(ipStr string) ([]byte, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, fmt.Errorf("Not a valid IP: %s", ipStr)
	}
	return []byte(ip.To4()), nil
}

func MacToBinary(macStr string) ([]byte, error) {
	mac, err := net.ParseMAC(macStr)
	if err != nil {
		return nil, err
	}
	return []byte(mac), nil
}

func UInt32ToBinary(i uint32, numBytes int) ([]byte, error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b[numBytes:], nil
}