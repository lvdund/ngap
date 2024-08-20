package ie

type PduSessionResourceSetupResponseTransfer struct {
	DlQosFlowPerTnlInformation                    *QosFlowPerTnlInformation
	AdditionalDlQosFlowPerTnlInformation          *QosFlowPerTnlInformationList
	SecurityResult                                *SecurityResult
	QosFlowFailedToSetupList                      *QosFlowListWithCause
	RedundantDlQosFlowPerTnlInformation           *QosFlowPerTnlInformation
	AdditionalRedundantDlQosFlowPerTnlInformation *QosFlowPerTnlInformationList
	UsedRsnInformation                            *RedundantPduSessionInformation
	GlobalRanNodeIdOfSecondaryNgRanNode           *GlobalRanNodeId
	MbsSupportIndicator                           *MbsSupportIndicator
	MbsSessionSetupResponseList                   *MbsSessionSetupResponseList
	MbsSessionFailedToSetupList                   *MbsSessionFailedToSetupList
}
