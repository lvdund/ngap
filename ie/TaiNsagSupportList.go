package ie

import "ngap/aper"

type TaiNsagSupportList struct {
TaiNsagSupportItem	TaiNsagSupportItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiNsagSupportItem struct {
NsagId	aper.Integer	//`Integer:"valueLB:0,valueUB:255"`
NsagSliceSupportList	ExtendedSliceSupportList	//`bitstring:"sizeLB:0,sizeUB:150"`
}
