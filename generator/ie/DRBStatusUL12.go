package ie

import "gen/aper"

type DRBStatusUL12 struct {
	ULCOUNTValue              COUNTValueForPDCPSN12 `True,`
	ReceiveStatusOfULPDCPSDUs aper.BitString        `True,OPTIONAL`
	IEExtension               DRBStatusUL12ExtIEs   `False,OPTIONAL`
}
