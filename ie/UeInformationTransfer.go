package ie

type UeInformationTransfer struct {
	MessageType                  MessageType                  `bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GSTmsi                    Ie5GSTmsi                    `bitstring:"sizeLB:0,sizeUB:150"`
	NbIotUePriority              NbIotUePriority              `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapability            UeRadioCapability            `bitstring:"sizeLB:0,sizeUB:150"`
	SNssai                       SNssai                       `bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai                 AllowedNssai                 `bitstring:"sizeLB:0,sizeUB:150"`
	UeDifferentiationInformation UeDifferentiationInformation `bitstring:"sizeLB:0,sizeUB:150"`
	MaskedImeisv                 MaskedImeisv                 `bitstring:"sizeLB:0,sizeUB:150"`
}
