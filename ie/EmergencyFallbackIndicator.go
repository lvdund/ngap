package ie

type EmergencyFallbackIndicator struct {
EmergencyFallbackRequestIndicator	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
EmergencyServiceTargetCn	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
