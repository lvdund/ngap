package ie

type OverloadAction struct {
	OverloadAction []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
