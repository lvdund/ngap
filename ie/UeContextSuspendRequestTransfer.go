package ie

type UeContextSuspendRequestTransfer struct {
	SuspendIndicator []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
