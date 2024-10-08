package ie

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        PriorityLevelARP                     `False,`
	PreemptionCapability    PreemptionCapability                 `False,`
	PreemptionVulnerability PreemptionVulnerability              `False,`
	IEExtensions            AllocationAndRetentionPriorityExtIEs `False,OPTIONAL`
}
