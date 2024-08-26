package ie

type BroadcastCancelledAreaList struct {
	ChoiceBroadcastCancelledArea ChoiceBroadcastCancelledArea `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceBroadcastCancelledArea struct {
	CellIdCancelledEUtra          CellIdCancelledEUtra          `bitstring:"sizeLB:0,sizeUB:150"`
	TaiCancelledEUtra             TaiCancelledEUtra             `bitstring:"sizeLB:0,sizeUB:150"`
	EmergencyAreaIdCancelledEUtra EmergencyAreaIdCancelledEUtra `bitstring:"sizeLB:0,sizeUB:150"`
	CellIdCancelledNr             CellIdCancelledNr             `bitstring:"sizeLB:0,sizeUB:150"`
	TaiCancelledNr                TaiCancelledNr                `bitstring:"sizeLB:0,sizeUB:150"`
	EmergencyAreaIdCancelledNr    EmergencyAreaIdCancelledNr    `bitstring:"sizeLB:0,sizeUB:150"`
}

type CellIdCancelledEUtra struct {
	CancelledCellList CancelledCellList `bitstring:"sizeLB:0,sizeUB:150"`
}

type CancelledCellList struct {
	NrCgi              NrCgi              `bitstring:"sizeLB:0,sizeUB:150"`
	NumberOfBroadcasts NumberOfBroadcasts `bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiCancelledEUtra struct {
	TaiCancelled TaiCancelled `bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiCancelled struct {
	Tai                    Tai                    `bitstring:"sizeLB:0,sizeUB:150"`
	CancelledCellInTaiList CancelledCellInTaiList `bitstring:"sizeLB:0,sizeUB:150"`
}

type CancelledCellInTaiList struct {
	NrCgi              NrCgi              `bitstring:"sizeLB:0,sizeUB:150"`
	NumberOfBroadcasts NumberOfBroadcasts `bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdCancelledEUtra struct {
	EmergencyAreaIdCancelled EmergencyAreaIdCancelled `bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdCancelled struct {
	EmergencyAreaId                    EmergencyAreaId                    `bitstring:"sizeLB:0,sizeUB:150"`
	CancelledCellInEmergencyAreaIdList CancelledCellInEmergencyAreaIdList `bitstring:"sizeLB:0,sizeUB:150"`
}

type CancelledCellInEmergencyAreaIdList struct {
	NrCgi              NrCgi              `bitstring:"sizeLB:0,sizeUB:150"`
	NumberOfBroadcasts NumberOfBroadcasts `bitstring:"sizeLB:0,sizeUB:150"`
}

type CellIdCancelledNr struct {
	CancelledCellList CancelledCellList `bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiCancelledNr struct {
	TaiCancelled TaiCancelled `bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdCancelledNr struct {
	EmergencyAreaIdCancelled EmergencyAreaIdCancelled `bitstring:"sizeLB:0,sizeUB:150"`
}
