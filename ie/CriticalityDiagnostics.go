package ie

type CriticalityDiagnostics struct {
	ProcedureCode                            uint8 //`bitstring:"sizeLB:0,sizeUB:255"`
	TriggeringMessage                        *[]byte
	ProcedureCriticality                     *[]byte
	InformationElementCriticalityDiagnostics *InformationElementCriticalityDiagnostics
}

type InformationElementCriticalityDiagnostics struct {
	IeCriticality *[]byte
	IeId          uint16 //`bitstring:"sizeLB:0,sizeUB:65535"`
	TypeOfError   *[]byte
}
