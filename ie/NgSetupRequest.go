package ie

import (
	"fmt"
	"ngap/aper"
)

type NgSetupRequest struct {
	MessageType            MessageType            `bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeId        GlobalRanNodeId        `bitstring:"sizeLB:0,sizeUB:150"`
	RanNodeName            []byte                 `bitstring:"sizeLB:1,sizeUB:150"`
	SupportedTaList        []SupportedTaItem      `bitstring:"sizeLB:0,sizeUB:150"`
	DefaultPagingDrx       PagingDrx              `bitstring:"sizeLB:0,sizeUB:150"`
	UeRetentionInformation UeRetentionInformation `bitstring:"sizeLB:0,sizeUB:150"`
	NbIotDefaultPagingDrx  NbIotDefaultPagingDrx  `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedRanNodeName    ExtendedRanNodeName    `bitstring:"sizeLB:0,sizeUB:150"`
}

type SupportedTaItem struct {
	Tac                     Tac                     `bitstring:"sizeLB:0,sizeUB:150"`
	BroadcastPlmnList       []*BroadcastPlmnItem    `bitstring:"sizeLB:0,sizeUB:150"`
	ConfiguredTacIndication ConfiguredTacIndication `bitstring:"sizeLB:0,sizeUB:150"`
	RatInformation          RatInformation          `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *SupportedTaItem) Decode(r *aper.AperReader) (err error) {
	// var octets []byte
	// if octets, err = r.ReadOctetString(&aper.Constraint{
	// 	Lb: 3,
	// 	Ub: 3,
	// }, true); err != nil {
	// 	return
	// }

	// t.Tac = octets
	if err = ie.Tac.Decode(r); err != nil {
		return
	}
	if ie.BroadcastPlmnList, err = aper.ReadSequenceOfEx[*BroadcastPlmnItem](func() *BroadcastPlmnItem {
		return new(BroadcastPlmnItem)
	}, r, &aper.Constraint{
		Lb: 1,
		Ub: 12,
	}, true); err != nil {
		return
	}
	return
}

func (ie *SupportedTaItem) Encode(r *aper.AperWriter) (err error) {
	err = ie.Tac.Encode(r)
	return nil
}

type BroadcastPlmnItem struct {
	PlmnIdentity                PlmnIdentity             `bitstring:"sizeLB:0,sizeUB:150"`
	TaiSliceSupportList         SliceSupportList         `bitstring:"sizeLB:0,sizeUB:150"`
	NpnSupport                  NpnSupport               `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedTaiSliceSupportList ExtendedSliceSupportList `bitstring:"sizeLB:0,sizeUB:150"`
	TaiNsagSupportList          TaiNsagSupportList       `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *BroadcastPlmnItem) Decode(r *aper.AperReader) (err error) {
	fmt.Printf("decode Broadcast Plmn Item \n")
	if err = ie.PlmnIdentity.Decode(r); err != nil {
		return
	}
	return
}
