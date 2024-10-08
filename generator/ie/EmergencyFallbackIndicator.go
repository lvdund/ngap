package ie

type EmergencyFallbackIndicator struct {
	EmergencyFallbackRequestIndicator EmergencyFallbackRequestIndicator `False,`
	EmergencyServiceTargetCN          EmergencyServiceTargetCN          `False,OPTIONAL`
	IEExtensions                      EmergencyFallbackIndicatorExtIEs  `False,OPTIONAL`
}
