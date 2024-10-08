package ie

import "gen/aper"

type RecommendedCellItem struct {
	NGRANCGI         NGRANCGI                  `False,`
	TimeStayedInCell aper.Integer              `True,OPTIONAL`
	IEExtensions     RecommendedCellItemExtIEs `False,OPTIONAL`
}
