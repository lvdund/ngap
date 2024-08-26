package ie

type InterSystemResourceStatusReply struct {
	ReportingSystem ReportingSystem `bitstring:"sizeLB:0,sizeUB:150"`
}
