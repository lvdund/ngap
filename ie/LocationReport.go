package ie

type LocationReport struct {
	MessageType                    MessageType                    `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                    AmfUeNgapId                    `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                    RanUeNgapId                    `bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation        UserLocationInformation        `bitstring:"sizeLB:0,sizeUB:150"`
	UePresenceInAreaOfInterestList UePresenceInAreaOfInterestList `bitstring:"sizeLB:0,sizeUB:150"`
	LocationReportingRequestType   LocationReportingRequestType   `bitstring:"sizeLB:0,sizeUB:150"`
}
