package utils

import (
	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

type GlobalRanNodeId struct {
	PlmnId  *PlmnId
	N3IwfId string
	GNbId   *GNbId
	NgeNbId string
	WagfId  string
	TngfId  string
	TwifId  string
	Nid     string
	ENbId   string
}

type GNbId struct {
	BitLength int32
	GNBValue  string
}

func RanIdToModels(ranNodeId ies.GlobalRANNodeID) (ranId GlobalRanNodeId) {
	present := ranNodeId.Choice
	switch present {
	case ies.GlobalRANNodeIDPresentGlobalgnbId:
		ranId.GNbId = new(GNbId)
		gnbId := ranId.GNbId
		ngapGnbId := ranNodeId.GlobalGNBID
		plmnid := PlmnIdToModels(ngapGnbId.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapGnbId.GNBID.Choice == ies.GNBIDPresentGnbId {
			choiceGnbId := ngapGnbId.GNBID.GNBID
			gnbId.BitLength = int32(choiceGnbId.NumBits)
			gnbId.GNBValue = BitStringToHex(&aper.BitString{
				Bytes:   choiceGnbId.Bytes,
				NumBits: uint64(gnbId.BitLength),
			})
		}
	case ies.GlobalRANNodeIDPresentGlobalngenbId:
		ngapNgENBID := ranNodeId.GlobalNgENBID
		plmnid := PlmnIdToModels(ngapNgENBID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapNgENBID.NgENBID.Choice == ies.NgENBIDPresentMacrongenbId {
			macroNgENBID := ngapNgENBID.NgENBID.MacroNgENBID
			ranId.NgeNbId = "MacroNGeNB-" + BitStringToHex(&aper.BitString{
				Bytes:   macroNgENBID.Bytes,
				NumBits: uint64(macroNgENBID.NumBits),
			})
		} else if ngapNgENBID.NgENBID.Choice == ies.NgENBIDPresentShortmacrongenbId {
			shortMacroNgENBID := ngapNgENBID.NgENBID.ShortMacroNgENBID
			ranId.NgeNbId = "SMacroNGeNB-" + BitStringToHex(&aper.BitString{
				Bytes:   shortMacroNgENBID.Bytes,
				NumBits: uint64(shortMacroNgENBID.NumBits),
			})
		} else if ngapNgENBID.NgENBID.Choice == ies.NgENBIDPresentLongmacrongenbId {
			longMacroNgENBID := ngapNgENBID.NgENBID.LongMacroNgENBID
			ranId.NgeNbId = "LMacroNGeNB-" + BitStringToHex(&aper.BitString{
				Bytes:   longMacroNgENBID.Bytes,
				NumBits: uint64(longMacroNgENBID.NumBits),
			})
		}
	case ies.GlobalRANNodeIDPresentGlobaln3IwfId:
		ngapN3IWFID := ranNodeId.GlobalN3IWFID
		plmnid := PlmnIdToModels(ngapN3IWFID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapN3IWFID.N3IWFID.Choice == ies.N3IWFIDPresentN3IwfId {
			choiceN3IWFID := ngapN3IWFID.N3IWFID.N3IWFID
			ranId.N3IwfId = BitStringToHex(&aper.BitString{
				Bytes:   choiceN3IWFID.Bytes,
				NumBits: uint64(choiceN3IWFID.NumBits),
			})
		}
	case ies.GlobalRANNodeIDPresentChoiceExtensions:
	}

	return ranId
}

func RanIDToNgap(modelsRanNodeId GlobalRanNodeId) ies.GlobalRANNodeID {
	var ngapRanNodeId ies.GlobalRANNodeID

	if modelsRanNodeId.GNbId.BitLength != 0 {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobalgnbId
		ngapRanNodeId.GlobalGNBID = new(ies.GlobalGNBID)
		globalGNBID := ngapRanNodeId.GlobalGNBID

		globalGNBID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalGNBID.GNBID.Choice = ies.GNBIDPresentGnbId
		tmp := HexToBitString(modelsRanNodeId.GNbId.GNBValue, int(modelsRanNodeId.GNbId.BitLength))
		globalGNBID.GNBID.GNBID = &tmp
	} else if modelsRanNodeId.NgeNbId != "" {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobalngenbId
		ngapRanNodeId.GlobalNgENBID = new(ies.GlobalNgENBID)
		globalNgENBID := ngapRanNodeId.GlobalNgENBID

		globalNgENBID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		ngENBID := &globalNgENBID.NgENBID
		if modelsRanNodeId.NgeNbId[:11] == "MacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentMacrongenbId
			tmp := HexToBitString(modelsRanNodeId.NgeNbId[11:], 18)
			ngENBID.MacroNgENBID = &tmp
		} else if modelsRanNodeId.NgeNbId[:12] == "SMacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentShortmacrongenbId
			tmp := HexToBitString(modelsRanNodeId.NgeNbId[12:], 20)
			ngENBID.ShortMacroNgENBID = &tmp
		} else if modelsRanNodeId.NgeNbId[:12] == "LMacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentLongmacrongenbId
			tmp := HexToBitString(modelsRanNodeId.NgeNbId[12:], 21)
			ngENBID.LongMacroNgENBID = &tmp
		}
	} else if modelsRanNodeId.N3IwfId != "" {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobaln3IwfId
		ngapRanNodeId.GlobalN3IWFID = new(ies.GlobalN3IWFID)
		globalN3IWFID := ngapRanNodeId.GlobalN3IWFID

		globalN3IWFID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalN3IWFID.N3IWFID.Choice = ies.N3IWFIDPresentN3IwfId
		tmp := HexToBitString(modelsRanNodeId.N3IwfId, len(modelsRanNodeId.N3IwfId)*4)
		globalN3IWFID.N3IWFID.N3IWFID = &tmp
	}

	return ngapRanNodeId
}
