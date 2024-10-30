package utils

import (
	"github.com/lvdund/ngap/ies"
)

type AllowedSnssai struct {
	AllowedSnssai      *Snssai          `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
}
type NsiInformation struct {
	NrfId string `json:"nrfId" yaml:"nrfId"`
	NsiId string `json:"nsiId,omitempty" yaml:"nsiId"`
}

func AllowedNssaiToNgap(allowedNssaiModels []AllowedSnssai) (allowedNssaiNgap ies.AllowedNSSAI) {
	for _, allowedSnssai := range allowedNssaiModels {
		snssai := SNssaiToNgap(*allowedSnssai.AllowedSnssai)
		item := ies.AllowedNSSAIItem{
			SNSSAI: &snssai,
		}
		allowedNssaiNgap.Value = append(allowedNssaiNgap.Value, &item)
	}
	return
}

func AllowedNssaiToModels(allowedNssaiNgap ies.AllowedNSSAI) (allowedNssaiModels []AllowedSnssai) {
	for _, item := range allowedNssaiNgap.Value {
		snssai := SNssaiToModels(*item.SNSSAI)
		allowedSnssai := AllowedSnssai{
			AllowedSnssai: &snssai,
		}
		allowedNssaiModels = append(allowedNssaiModels, allowedSnssai)
	}
	return
}
