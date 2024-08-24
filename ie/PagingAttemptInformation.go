package ie

import "ngap/aper"

type PagingAttemptInformation struct {
PagingAttemptCount	aper.Integer	//`Integer:"valueLB:1,valueUB:16"`
IntendedNumberOfPagingAttempts	aper.Integer	//`Integer:"valueLB:1,valueUB:16"`
NextPagingAreaScope	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
