package ie

type MbsSessionTnlInformationNgRan struct {
	ChoiceSessionType ChoiceSessionType `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceSessionType struct {
	LocationIndependent LocationIndependent `bitstring:"sizeLB:0,sizeUB:150"`
	LocationDependent   LocationDependent   `bitstring:"sizeLB:0,sizeUB:150"`
}

type LocationIndependent struct {
	SharedNgUUnicastTnlInformation UpTransportLayerInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type LocationDependent struct {
	MbsSessionTnlInformationNgRanList []MbsSessionTnlInformationNgRanItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionTnlInformationNgRanItem struct {
	MbsAreaSessionId               MbsAreaSessionId            `bitstring:"sizeLB:0,sizeUB:150"`
	SharedNgUUnicastTnlInformation UpTransportLayerInformation `bitstring:"sizeLB:0,sizeUB:150"`
}
