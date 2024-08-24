package ie

type ConnectionEstablishmentIndication struct {
	MessageType                  MessageType                  //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                  AmfUeNgapId                  //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                  RanUeNgapId                  //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapability            UeRadioCapability            //`bitstring:"sizeLB:0,sizeUB:150"`
	EndIndication                EndIndication                //`bitstring:"sizeLB:0,sizeUB:150"`
	SNssai                       SNssai                       //`bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai                 AllowedNssai                 //`bitstring:"sizeLB:0,sizeUB:150"`
	UeDifferentiationInformation UeDifferentiationInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	DlCpSecurityInformation      DlCpSecurityInformation      //`bitstring:"sizeLB:0,sizeUB:150"`
	NbIotUePriority              NbIotUePriority              //`bitstring:"sizeLB:0,sizeUB:150"`
	EnhancedCoverageRestriction  EnhancedCoverageRestriction  //`bitstring:"sizeLB:0,sizeUB:150"`
	CeModeBRestricted            CeModeBRestricted            //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityId          UeRadioCapabilityId          //`bitstring:"sizeLB:0,sizeUB:150"`
	MaskedImeisv                 MaskedImeisv                 //`bitstring:"sizeLB:0,sizeUB:150"`
	OldAmf                       AmfName                      //`bitstring:"sizeLB:0,sizeUB:150"`
}
