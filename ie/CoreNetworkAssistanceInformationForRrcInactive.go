package ie

type CoreNetworkAssistanceInformationForRrcInactive struct {
	UeIdentityIndexValue                 UeIdentityIndexValue            `bitstring:"sizeLB:0,sizeUB:150"`
	UeSpecificDrx                        PagingDrx                       `bitstring:"sizeLB:0,sizeUB:150"`
	PeriodicRegistrationUpdateTimer      PeriodicRegistrationUpdateTimer `bitstring:"sizeLB:0,sizeUB:150"`
	MicoModeIndication                   MicoModeIndication              `bitstring:"sizeLB:0,sizeUB:150"`
	TaiListForRrcInactive                TaiListForRrcInactive           `bitstring:"sizeLB:0,sizeUB:150"`
	ExpectedUeBehaviour                  ExpectedUeBehaviour             `bitstring:"sizeLB:0,sizeUB:150"`
	EUtraPagingEdrxInformation           EUtraPagingEdrxInformation      `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedUeIdentityIndexValue         ExtendedUeIdentityIndexValue    `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityForPaging           UeRadioCapabilityForPaging      `bitstring:"sizeLB:0,sizeUB:150"`
	MicoAllPlmn                          MicoAllPlmn                     `bitstring:"sizeLB:0,sizeUB:150"`
	NrPagingEdrxInformation              NrPagingEdrxInformation         `bitstring:"sizeLB:0,sizeUB:150"`
	PagingCauseIndicationForVoiceService []byte                          `bitstring:"sizeLB:0,sizeUB:150"`
	PeipsAssistanceInformation           PeipsAssistanceInformation      `bitstring:"sizeLB:0,sizeUB:150"`
	HashedUeIdentityIndexValue           HashedUeIdentityIndexValue      `bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForRrcInactive struct {
	TaiListForRrcInactiveItem TaiListForRrcInactiveItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForRrcInactiveItem struct {
	Tai Tai `bitstring:"sizeLB:0,sizeUB:150"`
}
