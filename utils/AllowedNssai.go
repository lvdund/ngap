package utils

import "github.com/lvdund/ngap/ies"

type AllowedSnssai struct {
	AllowedSnssai      *Snssai
	NsiInformationList []NsiInformation
	MappedHomeSnssai   *Snssai
}

type NsiInformation struct {
	NrfId string
	NsiId string
}

func AllowedNssaiToNgap(allowedNssaiModels []AllowedSnssai) (allowedNssaiNgap []ies.AllowedNSSAIItem) {
	for _, allowedSnssai := range allowedNssaiModels {
		item := ies.AllowedNSSAIItem{
			SNSSAI: SNssaiToNgap(*allowedSnssai.AllowedSnssai),
		}
		allowedNssaiNgap = append(allowedNssaiNgap, item)
	}
	return
}

func AllowedNssaiToModels(allowedNssaiNgap []ies.AllowedNSSAIItem) (allowedNssaiModels []AllowedSnssai) {
	for _, item := range allowedNssaiNgap {
		snssai := SNssaiToModels(item.SNSSAI)
		allowedSnssai := AllowedSnssai{
			AllowedSnssai: &snssai,
		}
		allowedNssaiModels = append(allowedNssaiModels, allowedSnssai)
	}
	return
}
