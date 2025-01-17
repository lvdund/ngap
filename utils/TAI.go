package utils

import (
	"encoding/hex"

	"github.com/lvdund/ngap/ies"
	"github.com/sirupsen/logrus"
)

type Tai struct {
	PlmnId *PlmnId
	Tac    string
	Nid    string
}

func TaiToModels(tai ies.TAI) Tai {
	var modelsTai Tai

	plmnID := PlmnIdToModels(tai.PLMNIdentity)
	modelsTai.PlmnId = &plmnID
	modelsTai.Tac = hex.EncodeToString(tai.TAC)

	return modelsTai
}

func TaiToNgap(tai Tai) ies.TAI {
	var ngapTai ies.TAI

	ngapTai.PLMNIdentity = PlmnIdToNgap(*tai.PlmnId)
	if tac, err := hex.DecodeString(tai.Tac); err != nil {
		logrus.Warnf("Decode TAC failed: %+v", err)
	} else {
		ngapTai.TAC = tac
	}
	return ngapTai
}
