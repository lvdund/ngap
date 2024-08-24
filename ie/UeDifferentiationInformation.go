package ie

import "ngap/aper"

type UeDifferentiationInformation struct {
	PeriodicCommunicationIndicator []byte                     //`bitstring:"sizeLB:0,sizeUB:150"`
	PeriodicTime                   aper.Integer               //`Integer:"valueLB:1,valueUB:3600"`
	ScheduledCommunicationTime     ScheduledCommunicationTime //`bitstring:"sizeLB:0,sizeUB:150"`
	StationaryIndication           []byte                     //`bitstring:"sizeLB:0,sizeUB:150"`
	TrafficProfile                 []byte                     //`bitstring:"sizeLB:0,sizeUB:150"`
	BatteryIndication              []byte                     //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ScheduledCommunicationTime struct {
	DayOfWeek      aper.BitString //`bitstring:"sizeLB:7,sizeUB:7"`
	TimeOfDayStart aper.Integer   //`Integer:"valueLB:0,valueUB:86399"`
	TimeOfDayEnd   aper.Integer   //`Integer:"valueLB:0,valueUB:86399"`
}
