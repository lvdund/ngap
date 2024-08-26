package ie

type DownlinkNasTransport struct {
	MessageType                          MessageType                          `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                          AmfUeNgapId                          `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                          RanUeNgapId                          `bitstring:"sizeLB:0,sizeUB:150"`
	OldAmf                               AmfName                              `bitstring:"sizeLB:0,sizeUB:150"`
	RanPagingPriority                    RanPagingPriority                    `bitstring:"sizeLB:0,sizeUB:150"`
	NasPdu                               NasPdu                               `bitstring:"sizeLB:0,sizeUB:150"`
	MobilityRestrictionList              MobilityRestrictionList              `bitstring:"sizeLB:0,sizeUB:150"`
	IndexToRatFrequencySelectionPriority IndexToRatFrequencySelectionPriority `bitstring:"sizeLB:0,sizeUB:150"`
	UeAggregateMaximumBitRate            UeAggregateMaximumBitRate            `bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai                         AllowedNssai                         `bitstring:"sizeLB:0,sizeUB:150"`
	SrvccOperationPossible               SrvccOperationPossible               `bitstring:"sizeLB:0,sizeUB:150"`
	EnhancedCoverageRestriction          EnhancedCoverageRestriction          `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedConnectedTime                ExtendedConnectedTime                `bitstring:"sizeLB:0,sizeUB:150"`
	UeDifferentiationInformation         UeDifferentiationInformation         `bitstring:"sizeLB:0,sizeUB:150"`
	CeModeBRestricted                    CeModeBRestricted                    `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapability                    UeRadioCapability                    `bitstring:"sizeLB:0,sizeUB:150"`
	UeCapabilityInfoRequest              UeCapabilityInfoRequest              `bitstring:"sizeLB:0,sizeUB:150"`
	EndIndication                        EndIndication                        `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityId                  UeRadioCapabilityId                  `bitstring:"sizeLB:0,sizeUB:150"`
	TargetNssaiInformation               TargetNssaiInformation               `bitstring:"sizeLB:0,sizeUB:150"`
	MaskedImeisv                         MaskedImeisv                         `bitstring:"sizeLB:0,sizeUB:150"`
}
