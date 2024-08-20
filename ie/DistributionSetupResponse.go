package ie

type DistributionSetupResponse struct {
	MessageType                          *MessageType
	MbsSessionId                         *MbsSessionId
	MbsAreaSessionId                     *MbsAreaSessionId
	MbsDistributionSetupResponseTransfer *[]byte
	CriticalityDiagnostics               *CriticalityDiagnostics
}
