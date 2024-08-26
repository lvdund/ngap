package ie

type NgRanCgi struct {
	ChoiceNgRanCgi ChoiceNgRanCgi `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceNgRanCgi struct {
	Nr    Nr    `bitstring:"sizeLB:0,sizeUB:150"`
	EUtra EUtra `bitstring:"sizeLB:0,sizeUB:150"`
}
