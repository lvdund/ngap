package ie

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          NGRANCGI                              `False,`
	CellType                              CellType                              `True,`
	TimeUEStayedInCell                    TimeUEStayedInCell                    `False,`
	TimeUEStayedInCellEnhancedGranularity TimeUEStayedInCellEnhancedGranularity `False,OPTIONAL`
	HOCauseValue                          Cause                                 `False,OPTIONAL`
	IEExtensions                          LastVisitedNGRANCellInformationExtIEs `False,OPTIONAL`
}
