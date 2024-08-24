package ie

type Ie5GProseAuthorized struct {
Ie5GProseDirectDiscovery	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProseDirectCommunication	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProseLayer2UeToNetworkRelay	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProseLayer3UeToNetworkRelay	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProseLayer2RemoteUe	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
