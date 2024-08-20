package ie

type AmfStatusIndication struct {
MessageType	*MessageType
UnavailableGuamiList	*[]UnavailableGuamiItem
}

type UnavailableGuamiItem struct {
Guami	*Guami
TimerApproachForGuamiRemoval	*[]byte
BackupAmfName	*AmfName
}
