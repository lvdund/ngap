package utils

import (
	"encoding/binary"

	"github.com/lvdund/ngap/aper"
	"github.com/sirupsen/logrus"
)

func TimeStampToInt32(timeStampNgap aper.OctetString) (timeStamp int32) {
	if len(timeStampNgap) != 4 {
		logrus.Error("TimeStampToInt32: the size of OctetString is not 4")
	}

	timeStamp = int32(binary.BigEndian.Uint32(timeStampNgap))
	return
}

func TimeStampToNgap(timeStamp int32) (timeStampNgap aper.OctetString) {
	// TODO: finish this function when need
	return
}
