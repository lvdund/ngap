package ie

import "gen/aper"

type COUNTValueForPDCPSN12 struct {
	PDCPSN12     aper.Integer                `False,`
	HFNPDCPSN12  aper.Integer                `False,`
	IEExtensions COUNTValueForPDCPSN12ExtIEs `False,OPTIONAL`
}
