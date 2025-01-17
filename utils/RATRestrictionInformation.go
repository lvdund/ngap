package utils

type RatType string

// List of RatType
const (
	RatType_NR             RatType = "NR"
	RatType_EUTRA          RatType = "EUTRA"
	RatType_WLAN           RatType = "WLAN"
	RatType_VIRTUAL        RatType = "VIRTUAL"
	RatType_NBIOT          RatType = "NBIOT"
	RatType_WIRELINE       RatType = "WIRELINE"
	RatType_WIRELINE_CABLE RatType = "WIRELINE_CABLE"
	RatType_WIRELINE_BBF   RatType = "WIRELINE_BBF"
	RatType_LTE_M          RatType = "LTE-M"
	RatType_NR_U           RatType = "NR_U"
	RatType_EUTRA_U        RatType = "EUTRA_U"
	RatType_TRUSTED_N3_GA  RatType = "TRUSTED_N3GA"
	RatType_TRUSTED_WLAN   RatType = "TRUSTED_WLAN"
	RatType_UTRA           RatType = "UTRA"
	RatType_GERA           RatType = "GERA"
)

// TS 38.413 9.3.1.85
func RATRestrictionInformationToNgap(ratType RatType) (ratResInfo []byte) {
	// Only support EUTRA & NR in version15.2.0
	switch ratType {
	case RatType_EUTRA:
		ratResInfo = []byte{0x80}
	case RatType_NR:
		ratResInfo = []byte{0x40}
	}
	return
}
