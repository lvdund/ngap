package utils

import (
	"net"

	"github.com/lvdund/ngap/aper"
	"github.com/sirupsen/logrus"
)

func IPAddressToString(ipAddr []byte) (ipv4Addr, ipv6Addr string) {

	// Described in 38.414
	switch len(ipAddr) * 8 {
	case 32: // ipv4
		netIP := net.IPv4(ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
		ipv4Addr = netIP.String()
	case 128: // ipv6
		netIP := net.IP{}
		for i := range ipAddr {
			netIP = append(netIP, ipAddr[i])
		}
		ipv6Addr = netIP.String()
	case 160: // ipv4 + ipv6, and ipv4 is contained in the first 32 bits
		netIPv4 := net.IPv4(ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
		netIPv6 := net.IP{}
		for i := range ipAddr {
			netIPv6 = append(netIPv6, ipAddr[i+4])
		}
		ipv4Addr = netIPv4.String()
		ipv6Addr = netIPv6.String()
	}
	return
}

func IPAddressToNgap(ipv4Addr, ipv6Addr string) (ipAddr aper.BitString) {

	if ipv4Addr == "" && ipv6Addr == "" {
		logrus.Warningln("IPAddressToNgap: Both ipv4 & ipv6 are nil string")
		return ipAddr
	}

	if ipv4Addr != "" && ipv6Addr != "" { // Both ipv4 & ipv6
		ipv4NetIP := net.ParseIP(ipv4Addr).To4()
		ipv6NetIP := net.ParseIP(ipv6Addr).To16()

		ipBytes := []byte{ipv4NetIP[0], ipv4NetIP[1], ipv4NetIP[2], ipv4NetIP[3]}
		for i := 0; i < 16; i++ {
			ipBytes = append(ipBytes, ipv6NetIP[i])
		}

		ipAddr = aper.BitString{
			Bytes: ipBytes,
			NumBits: 160,
		}

	} else if ipv4Addr != "" && ipv6Addr == "" { // ipv4
		ipv4NetIP := net.ParseIP(ipv4Addr).To4()

		ipBytes := []byte{ipv4NetIP[0], ipv4NetIP[1], ipv4NetIP[2], ipv4NetIP[3]}

		ipAddr = aper.BitString{
			Bytes: ipBytes,
			NumBits: 32,
		}

	} else { // ipv6
		ipv6NetIP := net.ParseIP(ipv6Addr).To16()

		ipBytes := []byte{}
		for i := 0; i < 16; i++ {
			ipBytes = append(ipBytes, ipv6NetIP[i])
		}

		ipAddr = aper.BitString{
			Bytes: ipBytes,
			NumBits: 128,
		}

	}

	return ipAddr
}
