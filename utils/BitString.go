package utils

import (
	"encoding/hex"

	"github.com/lvdund/ngap/aper"
	"github.com/sirupsen/logrus"
)

func BitStringToHex(bitString *aper.BitString) (hexString string) {
	hexString = hex.EncodeToString(bitString.Bytes)
	hexLen := (bitString.NumBits + 3) / 4
	hexString = hexString[:hexLen]
	return
}

func HexToBitString(hexString string, bitLength int) (bitString aper.BitString) {
	hexLen := len(hexString)
	if hexLen != (bitLength+3)/4 {
		logrus.Warningln("hexLen[", hexLen, "] doesn't match bitLength[", bitLength, "]")
		return
	}
	if hexLen%2 == 1 {
		hexString += "0"
	}
	if byteTmp, err := hex.DecodeString(hexString); err != nil {
		logrus.Warnf("Decode byteString failed: %+v", err)
	} else {
		bitString.Bytes = byteTmp
	}
	bitString.NumBits = uint64(bitLength)
	mask := byte(0xff)
	mask = mask << uint(8-bitLength%8)
	if mask != 0 {
		bitString.Bytes[len(bitString.Bytes)-1] &= mask
	}
	return
}

func ByteToBitString(byteArray []byte, bitLength int) (bitString aper.BitString) {
	byteLen := (bitLength + 7) / 8
	if byteLen > len(byteArray) {
		logrus.Warningln("bitLength[", bitLength, "] is beyond byteArray size[", len(byteArray), "]")
		return
	}
	bitString.Bytes = byteArray
	bitString.NumBits = uint64(bitLength)
	return
}
