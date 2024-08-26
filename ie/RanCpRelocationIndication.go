package ie

type RanCpRelocationIndication struct {
	MessageType             MessageType             `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId             RanUeNgapId             `bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GSTmsi               Ie5GSTmsi               `bitstring:"sizeLB:0,sizeUB:150"`
	EUtraCgi                EUtraCgi                `bitstring:"sizeLB:0,sizeUB:150"`
	Tai                     Tai                     `bitstring:"sizeLB:0,sizeUB:150"`
	UlCpSecurityInformation UlCpSecurityInformation `bitstring:"sizeLB:0,sizeUB:150"`
}
