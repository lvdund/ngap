package utils

import (
	"encoding/hex"

	"github.com/lvdund/ngap/ies"
)

type Tai struct {
	PlmnId *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	Tac    string  `json:"tac" yaml:"tac" bson:"tac" mapstructure:"Tac"`
	Nid    string  `json:"nid,omitempty" yaml:"nid" bson:"nid" mapstructure:"Nid"`
}

func TaiToModels(tai ies.TAI) Tai {
	var modelsTai Tai

	plmnID := PlmnIdToModels(*tai.PLMNIdentity)
	modelsTai.PlmnId = &plmnID
	modelsTai.Tac = hex.EncodeToString(tai.TAC.Value)

	return modelsTai
}

func TaiToNgap(tai Tai) ies.TAI {
	var ngapTai ies.TAI

	*ngapTai.PLMNIdentity = PlmnIdToNgap(*tai.PlmnId)
	if tac, err := hex.DecodeString(tai.Tac); err != nil {
		// logger.NgapLog.Warnf("Decode TAC failed: %+v", err)
	} else {
		ngapTai.TAC.Value = tac
	}
	return ngapTai
}
