package ie

import "gen/aper"

type COUNTValueForPDCPSN18 struct {
	PDCPSN18     aper.Integer                `False,`
	HFNPDCPSN18  aper.Integer                `False,`
	IEExtensions COUNTValueForPDCPSN18ExtIEs `False,OPTIONAL`
}
