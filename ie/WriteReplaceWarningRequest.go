package ie

import "ngap/aper"

type WriteReplaceWarningRequest struct {
	MessageType                       MessageType                       //`bitstring:"sizeLB:0,sizeUB:150"`
	MessageIdentifier                 MessageIdentifier                 //`bitstring:"sizeLB:0,sizeUB:150"`
	SerialNumber                      SerialNumber                      //`bitstring:"sizeLB:0,sizeUB:150"`
	WarningAreaList                   WarningAreaList                   //`bitstring:"sizeLB:0,sizeUB:150"`
	RepetitionPeriod                  RepetitionPeriod                  //`bitstring:"sizeLB:0,sizeUB:150"`
	NumberOfBroadcastsRequested       NumberOfBroadcastsRequested       //`bitstring:"sizeLB:0,sizeUB:150"`
	WarningType                       WarningType                       //`bitstring:"sizeLB:0,sizeUB:150"`
	WarningSecurityInformation        aper.OctetString                  //`octetstring:"sizeLB:50,sizeUB:50"`
	DataCodingScheme                  DataCodingScheme                  //`bitstring:"sizeLB:0,sizeUB:150"`
	WarningMessageContents            WarningMessageContents            //`bitstring:"sizeLB:0,sizeUB:150"`
	ConcurrentWarningMessageIndicator ConcurrentWarningMessageIndicator //`bitstring:"sizeLB:0,sizeUB:150"`
	WarningAreaCoordinates            WarningAreaCoordinates            //`bitstring:"sizeLB:0,sizeUB:150"`
}
