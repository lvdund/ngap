package ie

import "gen/aper"

type N3IWFID struct {
	N3IWFID          aper.BitString `False,`
	ChoiceExtensions N3IWFIDExtIEs  `False,`
}
