package ie

type HandoverCommand struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
HandoverType	*HandoverType
NasSecurityParametersFromNgRan	*NasSecurityParametersFromNgRan
PduSessionResourceHandoverList	*[]PduSessionResourceHandoverItem
PduSessionResourceToReleaseList	*[]PduSessionResourceToReleaseItem
TargetToSourceTransparentContainer	*TargetToSourceTransparentContainer
CriticalityDiagnostics	*CriticalityDiagnostics
}

type PduSessionResourceHandoverItem struct {
PduSessionId	*PduSessionId
HandoverCommandTransfer	*[]byte
}
