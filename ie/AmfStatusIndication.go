package ie

type AmfStatusIndication struct {
	MessageType          MessageType            `bitstring:"sizeLB:0,sizeUB:150"`
	UnavailableGuamiList []UnavailableGuamiItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type UnavailableGuamiItem struct {
	Guami                        Guami   `bitstring:"sizeLB:0,sizeUB:150"`
	TimerApproachForGuamiRemoval []byte  `bitstring:"sizeLB:0,sizeUB:150"`
	BackupAmfName                AmfName `bitstring:"sizeLB:0,sizeUB:150"`
}
