package ie

type MbsSessionFailedToSetupList struct {
	MbsSessionFailedToSetupList []MbsSessionFailedToSetupItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionFailedToSetupItem struct {
	MbsSessionId     MbsSessionId     //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId MbsAreaSessionId //`bitstring:"sizeLB:0,sizeUB:150"`
	Cause            Cause            //`bitstring:"sizeLB:0,sizeUB:150"`
}
