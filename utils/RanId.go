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
			gnbId.BitLength = int32(len(choiceGnbId) * 8)
			gnbId.GNBValue = BitStringToHex(&aper.BitString{
				Bytes:   choiceGnbId,
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
				Bytes:   macroNgENBID,
				NumBits: uint64(len(macroNgENBID) * 8),
			})
		} else if ngapNgENBID.NgENBID.Choice == ies.NgENBIDPresentShortmacrongenbId {
			shortMacroNgENBID := ngapNgENBID.NgENBID.ShortMacroNgENBID
			ranId.NgeNbId = "SMacroNGeNB-" + BitStringToHex(&aper.BitString{
				Bytes:   shortMacroNgENBID,
				NumBits: uint64(len(shortMacroNgENBID) * 8),
			})
		} else if ngapNgENBID.NgENBID.Choice == ies.NgENBIDPresentLongmacrongenbId {
			longMacroNgENBID := ngapNgENBID.NgENBID.LongMacroNgENBID
			ranId.NgeNbId = "LMacroNGeNB-" + BitStringToHex(&aper.BitString{
				Bytes:   longMacroNgENBID,
				NumBits: uint64(len(longMacroNgENBID) * 8),
			})
		}
	case ies.GlobalRANNodeIDPresentGlobaln3IwfId:
		ngapN3IWFID := ranNodeId.GlobalN3IWFID
		plmnid := PlmnIdToModels(ngapN3IWFID.PLMNIdentity)
		ranId.PlmnId = &plmnid
		if ngapN3IWFID.N3IWFID.Choice == ies.N3IWFIDPresentN3IwfId {
			choiceN3IWFID := ngapN3IWFID.N3IWFID.N3IWFID
			ranId.N3IwfId = BitStringToHex(&aper.BitString{
				Bytes:   choiceN3IWFID,
				NumBits: uint64(len(choiceN3IWFID) * 8),
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
		globalGNBID.GNBID.GNBID = HexToBitString(modelsRanNodeId.GNbId.GNBValue, int(modelsRanNodeId.GNbId.BitLength)).Bytes
	} else if modelsRanNodeId.NgeNbId != "" {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobalngenbId
		ngapRanNodeId.GlobalNgENBID = new(ies.GlobalNgENBID)
		globalNgENBID := ngapRanNodeId.GlobalNgENBID

		globalNgENBID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		ngENBID := &globalNgENBID.NgENBID
		if modelsRanNodeId.NgeNbId[:11] == "MacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentMacrongenbId
			ngENBID.MacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[11:], 18).Bytes
		} else if modelsRanNodeId.NgeNbId[:12] == "SMacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentShortmacrongenbId
			ngENBID.ShortMacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[12:], 20).Bytes
		} else if modelsRanNodeId.NgeNbId[:12] == "LMacroNGeNB-" {
			ngENBID.Choice = ies.NgENBIDPresentLongmacrongenbId
			ngENBID.LongMacroNgENBID = HexToBitString(modelsRanNodeId.NgeNbId[12:], 21).Bytes
		}
	} else if modelsRanNodeId.N3IwfId != "" {
		ngapRanNodeId.Choice = ies.GlobalRANNodeIDPresentGlobaln3IwfId
		ngapRanNodeId.GlobalN3IWFID = new(ies.GlobalN3IWFID)
		globalN3IWFID := ngapRanNodeId.GlobalN3IWFID

		globalN3IWFID.PLMNIdentity = PlmnIdToNgap(*modelsRanNodeId.PlmnId)
		globalN3IWFID.N3IWFID.Choice = ies.N3IWFIDPresentN3IwfId
		globalN3IWFID.N3IWFID.N3IWFID = HexToBitString(modelsRanNodeId.N3IwfId, len(modelsRanNodeId.N3IwfId)*4).Bytes
	}

	return ngapRanNodeId
}
