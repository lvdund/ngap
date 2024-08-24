package ie

type PduSessionResourceSetupResponseTransfer struct {
	DlQosFlowPerTnlInformation                    QosFlowPerTnlInformation       //`bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalDlQosFlowPerTnlInformation          QosFlowPerTnlInformationList   //`bitstring:"sizeLB:0,sizeUB:150"`
	SecurityResult                                SecurityResult                 //`bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowFailedToSetupList                      QosFlowListWithCause           //`bitstring:"sizeLB:0,sizeUB:150"`
	RedundantDlQosFlowPerTnlInformation           QosFlowPerTnlInformation       //`bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalRedundantDlQosFlowPerTnlInformation QosFlowPerTnlInformationList   //`bitstring:"sizeLB:0,sizeUB:150"`
	UsedRsnInformation                            RedundantPduSessionInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeIdOfSecondaryNgRanNode           GlobalRanNodeId                //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSupportIndicator                           MbsSupportIndicator            //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionSetupResponseList                   MbsSessionSetupResponseList    //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionFailedToSetupList                   MbsSessionFailedToSetupList    //`bitstring:"sizeLB:0,sizeUB:150"`
}
