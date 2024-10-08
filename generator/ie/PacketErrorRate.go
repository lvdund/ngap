package ie

import "gen/aper"

type PacketErrorRate struct {
	PERExponent  aper.Integer          `False,`
	IEExtensions PacketErrorRateExtIEs `False,OPTIONAL`
}
