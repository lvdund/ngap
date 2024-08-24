package ie

type UeHistoryInformation struct {
	LastVisitedCellItem LastVisitedCellItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}
