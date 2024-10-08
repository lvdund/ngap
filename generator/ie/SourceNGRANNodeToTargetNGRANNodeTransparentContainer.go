package ie

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      RRCContainer                                               `False,`
	PDUSessionResourceInformationList PDUSessionResourceInformationList                          `False,OPTIONAL`
	ERABInformationList               ERABInformationList                                        `False,OPTIONAL`
	TargetCellID                      NGRANCGI                                                   `False,`
	IndexToRFSP                       IndexToRFSP                                                `False,OPTIONAL`
	UEHistoryInformation              UEHistoryInformation                                       `False,`
	IEExtensions                      SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs `False,OPTIONAL`
}
