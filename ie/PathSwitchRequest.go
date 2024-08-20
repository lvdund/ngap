package ie

type PathSwitchRequest struct {
	MessageType                                  *MessageType
	RanUeNgapId                                  *RanUeNgapId
	SourceAmfUeNgapId                            *AmfUeNgapId
	UserLocationInformation                      *UserLocationInformation
	UeSecurityCapabilities                       *UeSecurityCapabilities
	PduSessionResourceToBeSwitchedInDownlinkList *[]PduSessionResourceToBeSwitchedInDownlinkItem
	PduSessionResourceFailedToSetupList          *[]PduSessionResourceFailedToSetupItem
	RrcResumeCause                               *RrcEstablishmentCause
	RedcapIndication                             *RedcapIndication
}

type PduSessionResourceToBeSwitchedInDownlinkItem struct {
	PduSessionId              *PduSessionId
	PathSwitchRequestTransfer *[]byte
}

type PduSessionResourceFailedToSetupItem struct {
	PduSessionId                         *PduSessionId
	PathSwitchRequestSetupFailedTransfer *[]byte
}
