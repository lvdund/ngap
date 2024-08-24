package ie

type MbsServiceAreaInformation struct {
	MbsServiceAreaCellList MbsServiceAreaCellList //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsServiceAreaTaiList  MbsServiceAreaTaiList  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsServiceAreaCellList struct {
	NrCgi NrCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsServiceAreaTaiList struct {
	Tai Tai //`bitstring:"sizeLB:0,sizeUB:150"`
}
