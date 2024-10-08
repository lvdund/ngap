package ie

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   SliceOverloadList              `False,`
	SliceOverloadResponse               OverloadResponse               `False,OPTIONAL`
	SliceTrafficLoadReductionIndication TrafficLoadReductionIndication `False,OPTIONAL`
	IEExtensions                        OverloadStartNSSAIItemExtIEs   `False,OPTIONAL`
}
