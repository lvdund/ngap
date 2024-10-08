package ie

type LastVisitedCellInformation struct {
	NGRANCell        LastVisitedNGRANCellInformation  `False,`
	EUTRANCell       LastVisitedEUTRANCellInformation `False,`
	UTRANCell        LastVisitedUTRANCellInformation  `False,`
	GERANCell        LastVisitedGERANCellInformation  `False,`
	ChoiceExtensions LastVisitedCellInformationExtIEs `False,`
}
