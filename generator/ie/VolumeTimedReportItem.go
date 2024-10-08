package ie

import "gen/aper"

type VolumeTimedReportItem struct {
	StartTimeStamp aper.OctetString            `False,`
	EndTimeStamp   aper.OctetString            `False,`
	UsageCountUL   aper.Integer                `False,`
	UsageCountDL   aper.Integer                `False,`
	IEExtensions   VolumeTimedReportItemExtIEs `False,OPTIONAL`
}
