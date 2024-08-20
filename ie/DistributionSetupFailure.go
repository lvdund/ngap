package ie

type DistributionSetupFailure struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsAreaSessionId	*MbsAreaSessionId
MbsDistributionSetupUnsuccessfulTransfer	*[]byte
Cause	*Cause
CriticalityDiagnostics	*CriticalityDiagnostics
}
