package ie

type RetrieveUeInformation struct {
	MessageType MessageType `bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GSTmsi   Ie5GSTmsi   `bitstring:"sizeLB:0,sizeUB:150"`
}
