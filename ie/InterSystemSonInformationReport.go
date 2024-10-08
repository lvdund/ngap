package ie

type InterSystemSonInformationReport struct {
	ChoiceSonInformationReport ChoiceSonInformationReport `bitstring:"sizeLB:0,sizeUB:150"`
}

type EnergySavingsIndication struct {
	InterSystemCellStateIndication InterSystemCellStateIndication `bitstring:"sizeLB:0,sizeUB:150"`
}

type ResourceStatusReport struct {
	InterSystemResourceStatusReport InterSystemResourceStatusReport `bitstring:"sizeLB:0,sizeUB:150"`
}
