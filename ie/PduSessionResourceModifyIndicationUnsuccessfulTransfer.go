package ie

type PduSessionResourceModifyIndicationUnsuccessfulTransfer struct {
	Cause Cause `bitstring:"sizeLB:0,sizeUB:150"`
}
