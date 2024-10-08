package ie

import "gen/aper"

type NgENBID struct {
	MacroNgENBID      aper.BitString `False,`
	ShortMacroNgENBID aper.BitString `False,`
	LongMacroNgENBID  aper.BitString `False,`
	ChoiceExtensions  NgENBIDExtIEs  `False,`
}
