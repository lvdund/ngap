package ie

type DistributionReleaseRequest struct {
	MessageType                           *MessageType
	MbsSessionId                          *MbsSessionId
	MbsAreaSessionId                      *MbsAreaSessionId
	MbsDistributionReleaseRequestTransfer *[]byte
	Cause                                 *Cause
}
