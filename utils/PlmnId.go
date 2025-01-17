package utils

import (
	"encoding/hex"
	"strings"

	"github.com/sirupsen/logrus"
)

type PlmnId struct {
	Mcc string
	Mnc string
}

func PlmnIdToModels(ngapPlmnId []byte) (modelsPlmnid PlmnId) {
	hexString := strings.Split(hex.EncodeToString(ngapPlmnId), "")
	modelsPlmnid.Mcc = hexString[1] + hexString[0] + hexString[3]
	if hexString[2] == "f" {
		modelsPlmnid.Mnc = hexString[5] + hexString[4]
	} else {
		modelsPlmnid.Mnc = hexString[2] + hexString[5] + hexString[4]
	}
	return
}

func PlmnIdToNgap(modelsPlmnid PlmnId) (ngapPlmnId []byte) {
	var hexString string
	mcc := strings.Split(modelsPlmnid.Mcc, "")
	mnc := strings.Split(modelsPlmnid.Mnc, "")
	if len(modelsPlmnid.Mnc) == 2 {
		hexString = mcc[1] + mcc[0] + "f" + mcc[2] + mnc[1] + mnc[0]
	} else {
		hexString = mcc[1] + mcc[0] + mnc[0] + mcc[2] + mnc[2] + mnc[1]
	}

	if plmnId, err := hex.DecodeString(hexString); err != nil {
		logrus.Warnf("PlmnIdToNgap: Decode plmn failed: %+v", err)
	} else {
		ngapPlmnId = plmnId
	}
	return ngapPlmnId
}
