package ie

type LastVisitedCellInformation struct {
	ChoiceLastVisitedCellInformation *ChoiceLastVisitedCellInformation
}

type ChoiceLastVisitedCellInformation struct {
	NgRanCell  *NgRanCell
	EUtranCell *EUtranCell
	UtranCell  *UtranCell
	GeranCell  *GeranCell
}

type NgRanCell struct {
	LastVisitedNgRanCellInformation *LastVisitedNgRanCellInformation
}

type EUtranCell struct {
	LastVisitedEUtranCellInformation *[]byte
}

type UtranCell struct {
	LastVisitedUtranCellInformation *[]byte
}

type GeranCell struct {
	LastVisitedGeranCellInformation *[]byte
}
