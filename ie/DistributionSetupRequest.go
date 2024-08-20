package ie

type DistributionSetupRequest struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsAreaSessionId	*MbsAreaSessionId
MbsDistributionSetupRequestTransfer	*[]byte
}
