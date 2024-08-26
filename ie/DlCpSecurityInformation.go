package ie

import "ngap/aper"

type DlCpSecurityInformation struct {
	DlNasMac aper.BitString `bitstring:"sizeLB:16,sizeUB:16"`
}
