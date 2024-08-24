package ie

type EUtranCompositeAvailableCapacity struct {
	EUtranCellCapacityClassValue EUtranCellCapacityClassValue //`bitstring:"sizeLB:0,sizeUB:150"`
	EUtranCapacityValue          EUtranCapacityValue          //`bitstring:"sizeLB:0,sizeUB:150"`
}
