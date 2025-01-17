package utils

import (
	"encoding/hex"

	"github.com/lvdund/ngap/ies"
	"github.com/sirupsen/logrus"
)

type Snssai struct {
	Sst int32
	Sd  string
}

func SNssaiToModels(ngapSnssai ies.SNSSAI) (modelsSnssai Snssai) {
	modelsSnssai.Sst = int32(ngapSnssai.SST[0])
	if ngapSnssai.SD != nil {
		modelsSnssai.Sd = hex.EncodeToString(ngapSnssai.SD)
	}
	return
}

func SNssaiToNgap(modelsSnssai Snssai) ies.SNSSAI {
	var ngapSnssai ies.SNSSAI
	ngapSnssai.SST = []byte{byte(modelsSnssai.Sst)}

	if modelsSnssai.Sd != "" {
		if sdTmp, err := hex.DecodeString(modelsSnssai.Sd); err != nil {
			logrus.Warnf("Decode snssai.sd failed: %+v", err)
		} else {
			ngapSnssai.SD = sdTmp
		}
	}
	return ngapSnssai
}
