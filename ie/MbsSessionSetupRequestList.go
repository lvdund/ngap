package ie

type MbsSessionSetupRequestList struct {
	MbsSessionSetupRequestList *[]MbsSessionSetupRequestItem
}

type MbsSessionSetupRequestItem struct {
	MbsSessionId                         *MbsSessionId
	MbsAreaSessionId                     *MbsAreaSessionId
	AssociatedMbsQosFlowSetupRequestList *[]AssociatedMbsQosFlowSetupRequestItem
}

type AssociatedMbsQosFlowSetupRequestItem struct {
	MbsQosFlowIdentifier               *QosFlowIdentifier
	AssociatedUnicastQosFlowIdentifier *QosFlowIdentifier
}
