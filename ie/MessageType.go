package ie

type MessageType struct {
	ProcedureCode uint8 //`bitstring:"sizeLB:0,sizeUB:255"`
	TypeOfMessage *[]byte
}
