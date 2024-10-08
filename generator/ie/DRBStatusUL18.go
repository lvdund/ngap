package ie

import "gen/aper"

type DRBStatusUL18 struct {
	ULCOUNTValue              COUNTValueForPDCPSN18 `True,`
	ReceiveStatusOfULPDCPSDUs aper.BitString        `True,OPTIONAL`
	IEExtension               DRBStatusUL18ExtIEs   `False,OPTIONAL`
}
