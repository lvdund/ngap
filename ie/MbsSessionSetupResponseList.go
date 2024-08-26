package ie

type MbsSessionSetupResponseList struct {
	MbsSessionSetupResponseList []MbsSessionSetupResponseItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionSetupResponseItem struct {
	MbsSessionId     MbsSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId MbsAreaSessionId `bitstring:"sizeLB:0,sizeUB:150"`
}
