package utils

import (
	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

type GlobalRanNodeId struct {
	PlmnId  *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	N3IwfId string  `json:"n3IwfId,omitempty" yaml:"n3IwfId" bson:"n3IwfId" mapstructure:"N3IwfId"`
	GNbId   *GNbId  `json:"gNbId,omitempty" yaml:"gNbId" bson:"gNbId" mapstructure:"GNbId"`
	NgeNbId string  `json:"ngeNbId,omitempty" yaml:"ngeNbId" bson:"ngeNbId" mapstructure:"NgeNbId"`
	WagfId  string  `json:"wagfId,omitempty" yaml:"wagfId" bson:"wagfId" mapstructure:"WagfId"`
	TngfId  string  `json:"tngfId,omitempty" yaml:"tngfId" bson:"tngfId" mapstructure:"TngfId"`
	TwifId  string  `json:"twifId,omitempty" yaml:"twifId" bson:"twifId" mapstructure:"TwifId"`
	Nid     string  `json:"nid,omitempty" yaml:"nid" bson:"nid" mapstructure:"Nid"`
	ENbId   string  `json:"eNbId,omitempty" yaml:"eNbId" bson:"eNbId" mapstructure:"ENbId"`
}

type GNbId struct {
	BitLength int32  `json:"bitLength" yaml:"bitLength" bson:"bitLength" mapstructure:"BitLength"`
	GNBValue  string `json:"gNBValue" yaml:"gNBValue" bson:"gNBValue" mapstructure:"GNBValue"`
}

func RanIdToModels(ranNodeId ies.GlobalRANNodeID) (ranId GlobalRanNodeId) {
	present := ranNodeId.Choice
	switch present {
	case ies.GlobalRANNodeIDPresentGlobalGNBID:
		ranId.GNbId = new(GNbId)
		gnbId := ranId.GNbId
		ngapGnbId := ranNodeId.GlobalGNBID
		plmnid := PlmnIdToModels(*ngapGnbId.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapGnbId.GNBID.Choice == 1 {
			choiceGnbId := ngapGnbId.GNBID.GNBID
			gnbId.BitLength = int32(choiceGnbId.NumBits)
			gnbId.GNBValue = BitStringToHex(choiceGnbId)
		}
	case ies.GlobalRANNodeIDPresentGlobalNgENBID:
		ngapNgENBID := ranNodeId.GlobalNgENBID
		plmnid := PlmnIdToModels(*ngapNgENBID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapNgENBID.NgENBID.Choice == 1 {
			macroNgENBID := ngapNgENBID.NgENBID.MacroNgENBID
			ranId.NgeNbId = "MacroNGeNB-" + BitStringToHex(macroNgENBID)
		} else if ngapNgENBID.NgENBID.Choice == 2 {
			shortMacroNgENBID := ngapNgENBID.NgENBID.ShortMacroNgENBID
			ranId.NgeNbId = "SMacroNGeNB-" + BitStringToHex(shortMacroNgENBID)
		} else if ngapNgENBID.NgENBID.Choice == 3 {
			longMacroNgENBID := ngapNgENBID.NgENBID.LongMacroNgENBID
			ranId.NgeNbId = "LMacroNGeNB-" + BitStringToHex(longMacroNgENBID)
		}
	case ies.GlobalRANNodeIDPresentGlobalN3IWFID:
		ngapN3IWFID := ranNodeId.GlobalN3IWFID
		plmnid := PlmnIdToModels(*ngapN3IWFID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapN3IWFID.N3IWFID.Choice == 1 {
			choiceN3IWFID := ngapN3IWFID.N3IWFID.N3IWFID
			ranId.N3IwfId = BitStringToHex(choiceN3IWFID)
		}
	case 4:
		// switch ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.Present {
		// case 1:
		// 	ngapTNGFID := ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.GlobalTNGFID
		// 	plmnid := PlmnIdToModels(ngapTNGFID.PLMNIdentity)
		// 	ranId.PlmnId = &plmnid
		// 	if ngapTNGFID.TNGFID.Present == ies.TNGFIDPresentTNGFID {
		// 		choiceTNGFID := ngapTNGFID.TNGFID.TNGFID
		// 		ranId.TngfId = BitStringToHex(choiceTNGFID)
		// 	}
		// case 2:
		// 	ngapTWIFID := ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.GlobalTWIFID
		// 	plmnid := PlmnIdToModels(ngapTWIFID.PLMNIdentity)
		// 	ranId.PlmnId = &plmnid
		// 	if ngapTWIFID.TWIFID.Present == ies.TWIFIDPresentTWIFID {
		// 		choiceTWIFID := ngapTWIFID.TWIFID.TWIFID
		// 		ranId.TwifId = BitStringToHex(choiceTWIFID)
		// 	}
		// case 3:
		// 	ngapWAGFID := ranNodeId.ChoiceExtensions.GlobalRANNodeIDExtIEs.Value.GlobalWAGFID
		// 	plmnid := PlmnIdToModels(ngapWAGFID.PLMNIdentity)
		// 	ranId.PlmnId = &plmnid
		// 	if ngapWAGFID.WAGFID.Present == ies.WAGFIDPresentWAGFID {
		// 		choiceWAGFID := ngapWAGFID.WAGFID.WAGFID
		// 		ranId.WagfId = BitStringToHex(choiceWAGFID)
		// 	}
		// }
	}

	return ranId
}

func RanIDToNgap(modelsRanNodeId GlobalRanNodeId) ies.GlobalRANNodeID {
	var ngapRanNodeId ies.GlobalRANNodeID

	if modelsRanNodeId.GNbId.BitLength != 0 {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobalGNBID
		ngapRanNodeId.GlobalGNBID = new(ies.GlobalGNBID)
		globalGNBID := ngapRanNodeId.GlobalGNBID

		plmnid := PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalGNBID.PLMNIdentity = &plmnid
		globalGNBID.GNBID.Choice = ies.GNBIDPresentGNBID
		globalGNBID.GNBID.GNBID = new(aper.BitString)
		*globalGNBID.GNBID.GNBID = HexToBitString(modelsRanNodeId.GNbId.GNBValue, int(modelsRanNodeId.GNbId.BitLength))
	} else if modelsRanNodeId.NgeNbId != "" {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobalNgENBID
		ngapRanNodeId.GlobalNgENBID = new(ies.GlobalNgENBID)
		globalNgENBID := ngapRanNodeId.GlobalNgENBID

		plmnid := PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalNgENBID.PLMNIdentity = &plmnid
		ngENBID := globalNgENBID.NgENBID
		if modelsRanNodeId.NgeNbId[:11] == "MacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentMacroNgENBID
			ngENBID.MacroNgENBID = new(aper.BitString)
			b := HexToBitString(modelsRanNodeId.NgeNbId[11:], 18)
			ngENBID.MacroNgENBID = &b
		} else if modelsRanNodeId.NgeNbId[:12] == "SMacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentShortMacroNgENBID
			ngENBID.ShortMacroNgENBID = new(aper.BitString)
			*ngENBID.ShortMacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[12:], 20)
		} else if modelsRanNodeId.NgeNbId[:12] == "LMacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentLongMacroNgENBID
			ngENBID.LongMacroNgENBID = new(aper.BitString)
			*ngENBID.LongMacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[12:], 21)
		}
	} else if modelsRanNodeId.N3IwfId != "" {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobalN3IWFID
		ngapRanNodeId.GlobalN3IWFID = new(ies.GlobalN3IWFID)
		globalN3IWFID := ngapRanNodeId.GlobalN3IWFID

		*globalN3IWFID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalN3IWFID.N3IWFID.Choice = ies.N3IWFIDPresentN3IWFID
		globalN3IWFID.N3IWFID.N3IWFID = new(aper.BitString)
		*globalN3IWFID.N3IWFID.N3IWFID = HexToBitString(modelsRanNodeId.N3IwfId, len(modelsRanNodeId.N3IwfId)*4)
	}

	return ngapRanNodeId
}
