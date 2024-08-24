package ie

type InitialUeMessage struct {
	MessageType                         MessageType                         //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                         RanUeNgapId                         //`bitstring:"sizeLB:0,sizeUB:150"`
	NasPdu                              NasPdu                              //`bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation             UserLocationInformation             //`bitstring:"sizeLB:0,sizeUB:150"`
	RrcEstablishmentCause               RrcEstablishmentCause               //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GSTmsi                           Ie5GSTmsi                           //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfSetId                            AmfSetId                            //`bitstring:"sizeLB:0,sizeUB:150"`
	UeContextRequest                    []byte                              //`bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai                        AllowedNssai                        //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceToTargetAmfInformationReroute SourceToTargetAmfInformationReroute //`bitstring:"sizeLB:0,sizeUB:150"`
	SelectedPlmnIdentity                PlmnIdentity                        //`bitstring:"sizeLB:0,sizeUB:150"`
	IabNodeIndication                   []byte                              //`bitstring:"sizeLB:0,sizeUB:150"`
	CeModeBSupportIndicator             CeModeBSupportIndicator             //`bitstring:"sizeLB:0,sizeUB:150"`
	LteMIndication                      LteMIndication                      //`bitstring:"sizeLB:0,sizeUB:150"`
	EdtSession                          []byte                              //`bitstring:"sizeLB:0,sizeUB:150"`
	AuthenticatedIndication             []byte                              //`bitstring:"sizeLB:0,sizeUB:150"`
	NpnAccessInformation                NpnAccessInformation                //`bitstring:"sizeLB:0,sizeUB:150"`
	RedcapIndication                    RedcapIndication                    //`bitstring:"sizeLB:0,sizeUB:150"`
}
