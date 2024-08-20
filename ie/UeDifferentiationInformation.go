package ie

type UeDifferentiationInformation struct {
PeriodicCommunicationIndicator	*[]byte
PeriodicTime	uint16	//`bitstring:"sizeLB:1,sizeUB:3600"`
ScheduledCommunicationTime	*ScheduledCommunicationTime
StationaryIndication	*[]byte
TrafficProfile	*[]byte
BatteryIndication	*[]byte
}

type ScheduledCommunicationTime struct {
DayOfWeek	[]byte	//`bitstring:"sizeLB:7,sizeUB:7"`
TimeOfDayStart	uint32	//`bitstring:"sizeLB:0,sizeUB:86399"`
TimeOfDayEnd	uint32	//`bitstring:"sizeLB:0,sizeUB:86399"`
}
