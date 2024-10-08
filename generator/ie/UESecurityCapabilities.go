package ie

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             NRencryptionAlgorithms             `False,`
	NRintegrityProtectionAlgorithms    NRintegrityProtectionAlgorithms    `False,`
	EUTRAencryptionAlgorithms          EUTRAencryptionAlgorithms          `False,`
	EUTRAintegrityProtectionAlgorithms EUTRAintegrityProtectionAlgorithms `False,`
	IEExtensions                       UESecurityCapabilitiesExtIEs       `False,OPTIONAL`
}
