package ie

import "ngap/aper"

type CriticalityDiagnostics struct {
	ProcedureCode                            aper.Integer                             //`Integer:"valueLB:0,valueUB:255"`
	TriggeringMessage                        []byte                                   //`bitstring:"sizeLB:0,sizeUB:150"`
	ProcedureCriticality                     []byte                                   //`bitstring:"sizeLB:0,sizeUB:150"`
	InformationElementCriticalityDiagnostics InformationElementCriticalityDiagnostics //`bitstring:"sizeLB:0,sizeUB:150"`
}

type InformationElementCriticalityDiagnostics struct {
	IeCriticality []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
	IeId          aper.Integer //`Integer:"valueLB:0,valueUB:65535"`
	TypeOfError   []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
}
