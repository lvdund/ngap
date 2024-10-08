package ie

type ServedGUAMIItem struct {
	GUAMI         GUAMI                 `True,`
	BackupAMFName AMFName               `False,OPTIONAL`
	IEExtensions  ServedGUAMIItemExtIEs `False,OPTIONAL`
}
