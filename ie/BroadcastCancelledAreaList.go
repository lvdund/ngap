package ie

type BroadcastCancelledAreaList struct {
	ChoiceBroadcastCancelledArea *ChoiceBroadcastCancelledArea
}

type ChoiceBroadcastCancelledArea struct {
	CellIdCancelledEUtra          *CellIdCancelledEUtra
	TaiCancelledEUtra             *TaiCancelledEUtra
	EmergencyAreaIdCancelledEUtra *EmergencyAreaIdCancelledEUtra
	CellIdCancelledNr             *CellIdCancelledNr
	TaiCancelledNr                *TaiCancelledNr
	EmergencyAreaIdCancelledNr    *EmergencyAreaIdCancelledNr
}

type CellIdCancelledEUtra struct {
	CancelledCellList *CancelledCellList
}

type CancelledCellList struct {
	NrCgi              *NrCgi
	NumberOfBroadcasts *NumberOfBroadcasts
}

type TaiCancelledEUtra struct {
	TaiCancelled *TaiCancelled
}

type TaiCancelled struct {
	Tai                    *Tai
	CancelledCellInTaiList *CancelledCellInTaiList
}

type CancelledCellInTaiList struct {
	NrCgi              *NrCgi
	NumberOfBroadcasts *NumberOfBroadcasts
}

type EmergencyAreaIdCancelledEUtra struct {
	EmergencyAreaIdCancelled *EmergencyAreaIdCancelled
}

type EmergencyAreaIdCancelled struct {
	EmergencyAreaId                    *EmergencyAreaId
	CancelledCellInEmergencyAreaIdList *CancelledCellInEmergencyAreaIdList
}

type CancelledCellInEmergencyAreaIdList struct {
	NrCgi              *NrCgi
	NumberOfBroadcasts *NumberOfBroadcasts
}

type CellIdCancelledNr struct {
	CancelledCellList *CancelledCellList
}

type TaiCancelledNr struct {
	TaiCancelled *TaiCancelled
}

type EmergencyAreaIdCancelledNr struct {
	EmergencyAreaIdCancelled *EmergencyAreaIdCancelled
}
