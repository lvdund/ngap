package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/lvdund/ngap/ies"
)

type PlmnId struct {
	Mcc string `json:"mcc" yaml:"mcc" bson:"mcc" mapstructure:"Mcc"`
	Mnc string `json:"mnc" yaml:"mnc" bson:"mnc" mapstructure:"Mnc"`
}

func MccMncToPlmnid(mcc, mnc string) ([]byte, error) {
	if len(mcc) != 3 || (len(mnc) != 2 && len(mnc) != 3) {
		return nil, fmt.Errorf("invalid MCC or MNC length")
	}

	// Convert MCC digits
	mccDigit1, err := strconv.ParseUint(string(mcc[0]), 10, 4)
	if err != nil {
		return nil, err
	}
	mccDigit2, err := strconv.ParseUint(string(mcc[1]), 10, 4)
	if err != nil {
		return nil, err
	}
	mccDigit3, err := strconv.ParseUint(string(mcc[2]), 10, 4)
	if err != nil {
		return nil, err
	}

	// Convert MNC digits
	mncDigit1, err := strconv.ParseUint(string(mnc[0]), 10, 4)
	if err != nil {
		return nil, err
	}
	mncDigit2, err := strconv.ParseUint(string(mnc[1]), 10, 4)
	if err != nil {
		return nil, err
	}
	mncDigit3 := uint64(0xf) // Set to 0xf if MNC has only 2 digits
	if len(mnc) == 3 {
		mncDigit3, err = strconv.ParseUint(string(mnc[2]), 10, 4)
		if err != nil {
			return nil, err
		}
	}

	// Construct the 3-byte PLMN ID
	plmnID := make([]byte, 3)
	plmnID[0] = byte((mccDigit2 << 4) | mccDigit1)
	plmnID[1] = byte((mncDigit3 << 4) | mccDigit3)
	plmnID[2] = byte((mncDigit2 << 4) | mncDigit1)

	return plmnID, nil
}

func PlmnidToMccMnc(id []byte) (mcc, mnc string) {
	//decode from string here
	mccDigit1 := id[0] & 0x0f
	mccDigit2 := (id[0] & 0xf0) >> 4
	mccDigit3 := (id[1] & 0x0f)

	mncDigit1 := (id[2] & 0x0f)
	mncDigit2 := (id[2] & 0xf0) >> 4
	mncDigit3 := (id[1] & 0xf0) >> 4

	plmnIdBytes := []byte{(mccDigit1 << 4) | mccDigit2, (mccDigit3 << 4) | mncDigit1, (mncDigit2 << 4) | mncDigit3}
	plmnId := hex.EncodeToString(plmnIdBytes)
	mcc = plmnId[0:3]
	if plmnId[5] == 'f' {
		mnc = plmnId[3:5]
	} else {
		mnc = plmnId[3:6]
	}
	return
}

func PlmnIdToModels(ngapPlmnId ies.PLMNIdentity) (modelsPlmnid PlmnId) {
	value := ngapPlmnId.Value
	hexString := strings.Split(hex.EncodeToString(value), "")
	modelsPlmnid.Mcc = hexString[1] + hexString[0] + hexString[3]
	if hexString[2] == "f" {
		modelsPlmnid.Mnc = hexString[5] + hexString[4]
	} else {
		modelsPlmnid.Mnc = hexString[2] + hexString[5] + hexString[4]
	}
	return
}

func PlmnIdToNgap(modelsPlmnid PlmnId) ies.PLMNIdentity {
	var hexString string
	mcc := strings.Split(modelsPlmnid.Mcc, "")
	mnc := strings.Split(modelsPlmnid.Mnc, "")
	if len(modelsPlmnid.Mnc) == 2 {
		hexString = mcc[1] + mcc[0] + "f" + mcc[2] + mnc[1] + mnc[0]
	} else {
		hexString = mcc[1] + mcc[0] + mnc[0] + mcc[2] + mnc[2] + mnc[1]
	}

	var ngapPlmnId ies.PLMNIdentity
	if plmnId, err := hex.DecodeString(hexString); err != nil {
		// logger.NgapLog.Warnf("Decode plmn failed: %+v", err)
	} else {
		ngapPlmnId.Value = plmnId
	}
	return ngapPlmnId
}
