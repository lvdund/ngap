package ie

import (
	"fmt"
	"ngap/aper"
)

type NgSetupRequest struct {
	MessageType            MessageType
	GlobalRanNodeId        GlobalRanNodeId
	RanNodeName            []byte
	SupportedTaList        []SupportedTaItem
	DefaultPagingDrx       PagingDrx
	UeRetentionInformation UeRetentionInformation
	NbIotDefaultPagingDrx  NbIotDefaultPagingDrx
	ExtendedRanNodeName    ExtendedRanNodeName
}
type SupportedTaItem struct {
	Tac                     Tac                      //mandatory
	BroadcastPlmnList       []*BroadcastPlmnItem     //mandatory
	ConfiguredTacIndication *ConfiguredTacIndication //optional
	RatInformation          *RatInformation          //optional
}

func (ie *SupportedTaItem) Decode(r *aper.AperReader) (err error) {
	//SupportedTaItem has an extension bit (see free5gc SupportTaList which has
	//a ValueEx tag indicating inner items are extensible)
	var exBit bool
	if exBit, err = r.ReadBool(); err != nil {
		return
	}
	fmt.Printf("exBit=%v\n", exBit)

	//The structure has two optional fields, we need to read two bits to know
	//if these optional fields are set
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	fmt.Printf("optionals=%.8b\n", optionals)

	if err = ie.Tac.Decode(r); err != nil {
		return
	}
	if ie.BroadcastPlmnList, err = aper.ReadSequenceOfEx[*BroadcastPlmnItem](func() *BroadcastPlmnItem {
		return new(BroadcastPlmnItem)
	}, r, &aper.Constraint{
		Lb: 1,
		Ub: 12,
	}, false); err != nil {
		return
	}

	//TODO: now read optional fields
	if aper.IsBitSet(optionals, 0) {
	}
	if aper.IsBitSet(optionals, 1) {
	}

	return
}

func (ie *SupportedTaItem) Encode(r *aper.AperWriter) (err error) {
	err = ie.Tac.Encode(r)
	return nil
}

type BroadcastPlmnItem struct {
	PlmnIdentity                PlmnIdentity
	TaiSliceSupportList         []*SliceSupportItem
	NpnSupport                  *NpnSupport
	ExtendedTaiSliceSupportList *ExtendedSliceSupportList
	TaiNsagSupportList          *TaiNsagSupportList
}

func (ie *BroadcastPlmnItem) Decode(r *aper.AperReader) (err error) {
	fmt.Printf("decode Broadcast Plmn Item \n")

	//BroadcastPlmnItem has an extension bit (see free5gc BroadcastPlmnList which has
	//a ValueEx tag indicating inner items are extensible)
	var exBit bool
	if exBit, err = r.ReadBool(); err != nil {
		return
	}
	fmt.Printf("exBit=%v\n", exBit)

	//The structure has two optional fields, we need to read three bits to know
	//if these optional fields are set
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}

	if err = ie.PlmnIdentity.Decode(r); err != nil {
		return
	}

	fmt.Printf("decode TaiSliceSupportList \n")
	if ie.TaiSliceSupportList, err = aper.ReadSequenceOfEx[*SliceSupportItem](func() *SliceSupportItem {
		return new(SliceSupportItem)
	}, r, &aper.Constraint{
		Lb: 1,
		Ub: 1024,
	}, false); err != nil {
		return
	} else {
		fmt.Printf("Number of slices: %d\n", len(ie.TaiSliceSupportList))
	}

	//TODO: other optional fields

	if aper.IsBitSet(optionals, 0) {
	}
	if aper.IsBitSet(optionals, 1) {
	}
	if aper.IsBitSet(optionals, 2) {
	}

	return
}
