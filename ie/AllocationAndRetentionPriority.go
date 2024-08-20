package ie

type AllocationAndRetentionPriority struct {
PriorityLevel	uint8	//`bitstring:"sizeLB:1,sizeUB:15"`
PreEmptionCapability	*[]byte
PreEmptionVulnerability	*[]byte
}
