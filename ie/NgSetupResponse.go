package ie

type NgSetupResponse struct {
	MessageType            MessageType            //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfName                AmfName                //`bitstring:"sizeLB:0,sizeUB:150"`
	ServedGuamiList        []ServedGuamiItem      //`bitstring:"sizeLB:0,sizeUB:150"`
	RelativeAmfCapacity    RelativeAmfCapacity    //`bitstring:"sizeLB:0,sizeUB:150"`
	PlmnSupportList        []PlmnSupportItem      //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics CriticalityDiagnostics //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRetentionInformation UeRetentionInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	IabSupported           []byte                 //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedAmfName        ExtendedAmfName        //`bitstring:"sizeLB:0,sizeUB:150"`
}
