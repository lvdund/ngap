package ie

type PDUSessionResourceInformationItem struct {
	PDUSessionID              PDUSessionID                            `False,`
	QosFlowInformationList    QosFlowInformationList                  `False,`
	DRBsToQosFlowsMappingList DRBsToQosFlowsMappingList               `False,OPTIONAL`
	IEExtensions              PDUSessionResourceInformationItemExtIEs `False,OPTIONAL`
}
