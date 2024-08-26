package ie

type CellType struct {
	CellSize []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
