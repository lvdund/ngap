package ie

import "gen/aper"

type UEIdentityIndexValue struct {
	IndexLength10    aper.BitString             `False,`
	ChoiceExtensions UEIdentityIndexValueExtIEs `False,`
}
