package utils

import (
	"encoding/hex"

	"github.com/lvdund/ngap/ies"
)

type Snssai struct {
	Sst int32  `json:"sst" yaml:"sst" bson:"sst" mapstructure:"Sst"`
	Sd  string `json:"sd,omitempty" yaml:"sd" bson:"sd" mapstructure:"Sd"`
}

func SNssaiToModels(ngapSnssai ies.SNSSAI) (modelsSnssai Snssai) {
	modelsSnssai.Sst = int32(ngapSnssai.SST.Value[0])
	if ngapSnssai.SD != nil {
		modelsSnssai.Sd = hex.EncodeToString(ngapSnssai.SD.Value)
	}
	return
}

func SNssaiToNgap(modelsSnssai Snssai) ies.SNSSAI {
	var ngapSnssai ies.SNSSAI
	ngapSnssai.SST.Value = []byte{byte(modelsSnssai.Sst)}

	if modelsSnssai.Sd != "" {
		ngapSnssai.SD = new(ies.SD)
		if sdTmp, err := hex.DecodeString(modelsSnssai.Sd); err != nil {
			// logger.NgapLog.Warnf("Decode snssai.sd failed: %+v", err)
		} else {
			ngapSnssai.SD.Value = sdTmp
		}
	}
	return ngapSnssai
}
