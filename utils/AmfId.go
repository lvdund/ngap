package utils

import (
	"encoding/hex"

	"github.com/lvdund/ngap/aper"
	"github.com/sirupsen/logrus"
)

func AmfIdToNgap(amfId string) (regionId, setId, ptrId aper.BitString) {
	regionId = HexToBitString(amfId[:2], 8)
	setId = HexToBitString(amfId[2:5], 10)
	tmpByte, err := hex.DecodeString(amfId[4:])
	if err != nil {
		logrus.Warningln("AmfId From Models To NGAP Error: ", err.Error())
		return
	}
	shiftByte, err := aper.GetBitString(tmpByte, 2, 6)
	if err != nil {
		logrus.Warningln("AmfId From Models To NGAP Error: ", err.Error())
		return
	}
	ptrId.NumBits = 6
	ptrId.Bytes = shiftByte
	return
}

func AmfIdToModels(regionId, setId, ptrId aper.BitString) (amfId string) {
	regionHex := BitStringToHex(&regionId)
	tmpByte := []byte{setId.Bytes[0], (setId.Bytes[1] & 0xc0) | (ptrId.Bytes[0] >> 2)}
	restHex := hex.EncodeToString(tmpByte)
	amfId = regionHex + restHex
	return
}
