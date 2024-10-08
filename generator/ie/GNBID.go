package ie

import "gen/aper"

type GNBID struct {
	GNBID            aper.BitString `False,`
	ChoiceExtensions GNBIDExtIEs    `False,`
}
