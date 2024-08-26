package ie

import "ngap/aper"

type PathSwitchRequest struct {
	MessageType                                  MessageType                                    `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                                  RanUeNgapId                                    `bitstring:"sizeLB:0,sizeUB:150"`
	SourceAmfUeNgapId                            AmfUeNgapId                                    `bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation                      UserLocationInformation                        `bitstring:"sizeLB:0,sizeUB:150"`
	UeSecurityCapabilities                       UeSecurityCapabilities                         `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceToBeSwitchedInDownlinkList []PduSessionResourceToBeSwitchedInDownlinkItem `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToSetupList          []PduSessionResourceFailedToSetupItem          `bitstring:"sizeLB:0,sizeUB:150"`
	RrcResumeCause                               RrcEstablishmentCause                          `bitstring:"sizeLB:0,sizeUB:150"`
	RedcapIndication                             RedcapIndication                               `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceToBeSwitchedInDownlinkItem struct {
	PduSessionId              PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	PathSwitchRequestTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceFailedToSetupItem struct {
	PduSessionId                         PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	PathSwitchRequestSetupFailedTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
