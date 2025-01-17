package utils

import (
	"encoding/binary"
)

func PortNumberToInt(port []byte) (portInt32 int32) {
	portInt32 = int32(binary.BigEndian.Uint16(port))
	return
}

func PortNumberToNgap(portInt32 int32) (port []byte) {
	port = make([]byte, 2)
	binary.BigEndian.PutUint16(port, uint16(portInt32))
	return
}
