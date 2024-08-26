package ie

import "ngap/aper"

type AllocationAndRetentionPriority struct {
	PriorityLevel           aper.Integer `Integer:"valueLB:1,valueUB:15"`
	PreEmptionCapability    []byte       `bitstring:"sizeLB:0,sizeUB:150"`
	PreEmptionVulnerability []byte       `bitstring:"sizeLB:0,sizeUB:150"`
}
