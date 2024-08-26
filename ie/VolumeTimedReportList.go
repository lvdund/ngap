package ie

import "ngap/aper"

type VolumeTimedReportList struct {
	VolumeTimedReportItem VolumeTimedReportItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type VolumeTimedReportItem struct {
	StartTimestamp aper.OctetString `octetstring:"sizeLB:4,sizeUB:4"`
	EndTimestamp   aper.OctetString `octetstring:"sizeLB:4,sizeUB:4"`
	UsageCountUl   aper.Integer     `Integer:"valueLB:0,valueUB:18446744073709551615"`
	UsageCountDl   aper.Integer     `Integer:"valueLB:0,valueUB:18446744073709551615"`
}
