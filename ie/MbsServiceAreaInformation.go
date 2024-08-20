package ie

type MbsServiceAreaInformation struct {
MbsServiceAreaCellList	*MbsServiceAreaCellList
MbsServiceAreaTaiList	*MbsServiceAreaTaiList
}

type MbsServiceAreaCellList struct {
NrCgi	*NrCgi
}

type MbsServiceAreaTaiList struct {
Tai	*Tai
}
