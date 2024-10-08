package ie

type SonConfigurationTransfer struct {
	TargetRanNodeIdSon     TargetRanNodeIdSon     `bitstring:"sizeLB:0,sizeUB:150"`
	SourceRanNodeId        SourceRanNodeId        `bitstring:"sizeLB:0,sizeUB:150"`
	SonInformation         SonInformation         `bitstring:"sizeLB:0,sizeUB:150"`
	XnTnlConfigurationInfo XnTnlConfigurationInfo `bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetRanNodeIdSon struct {
	GlobalRanNodeId GlobalRanNodeId `bitstring:"sizeLB:0,sizeUB:150"`
	SelectedTai     Tai             `bitstring:"sizeLB:0,sizeUB:150"`
	NrCgi           NrCgi           `bitstring:"sizeLB:0,sizeUB:150"`
}
