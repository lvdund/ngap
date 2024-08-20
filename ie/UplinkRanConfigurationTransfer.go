package ie

type UplinkRanConfigurationTransfer struct {
	MessageType                         *MessageType
	SonConfigurationTransfer            *SonConfigurationTransfer
	EnDcSonConfigurationTransfer        *[]byte
	InterSystemSonConfigurationTransfer *InterSystemSonConfigurationTransfer
}
