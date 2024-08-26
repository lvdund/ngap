package ie

type TargetNgRanNodeToSourceNgRanNodeFailureTransparentContainer struct {
	CellCagInformation                   CellCagInformation                   `bitstring:"sizeLB:0,sizeUB:150"`
	NgapIeSupportInformationResponseList NgapIeSupportInformationResponseList `bitstring:"sizeLB:0,sizeUB:150"`
}
