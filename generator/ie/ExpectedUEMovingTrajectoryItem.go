package ie

import "gen/aper"

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         NGRANCGI                             `False,`
	TimeStayedInCell aper.Integer                         `True,OPTIONAL`
	IEExtensions     ExpectedUEMovingTrajectoryItemExtIEs `False,OPTIONAL`
}
