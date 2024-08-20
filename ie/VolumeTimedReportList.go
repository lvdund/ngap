package ie

type VolumeTimedReportList struct {
VolumeTimedReportItem	*VolumeTimedReportItem
}

type VolumeTimedReportItem struct {
StartTimestamp	[]byte	//`bitstring:"sizeLB:4,sizeUB:4"`
EndTimestamp	[]byte	//`bitstring:"sizeLB:4,sizeUB:4"`
UsageCountUl	uint64	//`bitstring:"sizeLB:0,sizeUB:18446744073709551615"`
UsageCountDl	uint64	//`bitstring:"sizeLB:0,sizeUB:18446744073709551615"`
}
