package ie

import "ngap/aper"

type MessageType struct {
	ProcedureCode aper.Integer //`Integer:"valueLB:0,valueUB:255"`
	TypeOfMessage []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
}
