package ie

type MbsSessionSetupRequestList struct {
	MbsSessionSetupRequestList []MbsSessionSetupRequestItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionSetupRequestItem struct {
	MbsSessionId                         MbsSessionId                           `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId                     MbsAreaSessionId                       `bitstring:"sizeLB:0,sizeUB:150"`
	AssociatedMbsQosFlowSetupRequestList []AssociatedMbsQosFlowSetupRequestItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type AssociatedMbsQosFlowSetupRequestItem struct {
	MbsQosFlowIdentifier               QosFlowIdentifier `bitstring:"sizeLB:0,sizeUB:150"`
	AssociatedUnicastQosFlowIdentifier QosFlowIdentifier `bitstring:"sizeLB:0,sizeUB:150"`
}
