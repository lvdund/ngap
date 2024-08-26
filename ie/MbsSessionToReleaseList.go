package ie

type MbsSessionToReleaseList struct {
	MbsSessionToReleaseList []MbsSessionToReleaseItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionToReleaseItem struct {
	MbsSessionId MbsSessionId `bitstring:"sizeLB:0,sizeUB:150"`
	Cause        Cause        `bitstring:"sizeLB:0,sizeUB:150"`
}
