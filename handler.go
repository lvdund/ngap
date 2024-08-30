package ngap

//This is to describe how NGAP handler should be written
func HandleNgap(wire []byte) {
	if pdu, err, diagnostics := NgapDecode(wire); err != nil {
		return
	} else {
		if diagnostics != nil {
			//TODO: send error indication
		}
	}

	switch pdu.Present {
	case NgapPduInitiatingMessage:
		handleInitiatingMessage(&pdu)
	case NgapPduSuccessfulOutcome:
		handleSuccessfulOutcome(&pdu)
	case NgapPduUnsuccessfulOutcome:
	}
}

func handleInitiatingMessage(pdu *NgapPdu) {
	switch pdu.Message.ProcedureCode {
	case NgapProcedureNGSetup: //NGSetupRequest
		msg := pdu.Message.Value.(*NGSetupRequest)
		//then handle the message
		handleNGSetupRequest(msg)
	default:
		//TODO:
	}
}

func handleSuccessfulOutcome(pdu *NgapPdu) {
	switch pdu.Message.ProcedureCode {
	case NgapProcedureNGSetup: //NGSetupResponse
		msg := pdu.Message.Value(*NGSetupResponse)
		//then handle the message
		handleNGSetupResponse(msg)
	default:
		//TODO:
	}
}

func handleNGSetupRequest(msg *NGSetupRequest) {
	//TODO: add example code
}

func handleNGSetupResponse(msg *NGSetupResponse) {
	//TODO: add example code
}
