package ie

import "ngap/aper"

type DownlinkRanConfigurationTransfer struct {
	MessageType                         MessageType                         //`bitstring:"sizeLB:0,sizeUB:150"`
	SonConfigurationTransfer            SonConfigurationTransfer            //`bitstring:"sizeLB:0,sizeUB:150"`
	EnDcSonConfigurationTransfer        aper.OctetString                    //`octetstring:"sizeLB:0,sizeUB:150"`
	InterSystemSonConfigurationTransfer InterSystemSonConfigurationTransfer //`bitstring:"sizeLB:0,sizeUB:150"`
}
