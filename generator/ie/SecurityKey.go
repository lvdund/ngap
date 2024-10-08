package ie

import "gen/aper"

type SecurityKey struct {
	Value aper.BitString `False,256,256`
}
