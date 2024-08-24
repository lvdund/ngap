package ie

import "ngap/aper"

type UeSecurityCapabilities struct {
NrEncryptionAlgorithms	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
NrIntegrityProtectionAlgorithms	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
EUtraEncryptionAlgorithms	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
EUtraIntegrityProtectionAlgorithms	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
}
