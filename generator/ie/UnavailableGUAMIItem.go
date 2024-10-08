package ie

type UnavailableGUAMIItem struct {
	GUAMI                        GUAMI                        `True,`
	TimerApproachForGUAMIRemoval TimerApproachForGUAMIRemoval `False,OPTIONAL`
	BackupAMFName                AMFName                      `False,OPTIONAL`
	IEExtensions                 UnavailableGUAMIItemExtIEs   `False,OPTIONAL`
}
