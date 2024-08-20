package ie

type InitialUeMessage struct {
	MessageType                         *MessageType
	RanUeNgapId                         *RanUeNgapId
	NasPdu                              *NasPdu
	UserLocationInformation             *UserLocationInformation
	RrcEstablishmentCause               *RrcEstablishmentCause
	Ie5GSTmsi                           *Ie5GSTmsi
	AmfSetId                            *AmfSetId
	UeContextRequest                    *[]byte
	AllowedNssai                        *AllowedNssai
	SourceToTargetAmfInformationReroute *SourceToTargetAmfInformationReroute
	SelectedPlmnIdentity                *PlmnIdentity
	IabNodeIndication                   *[]byte
	CeModeBSupportIndicator             *CeModeBSupportIndicator
	LteMIndication                      *LteMIndication
	EdtSession                          *[]byte
	AuthenticatedIndication             *[]byte
	NpnAccessInformation                *NpnAccessInformation
	RedcapIndication                    *RedcapIndication
}
