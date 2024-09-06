package ngap

//hold a decoded Ngap message
type NgapPdu struct {
	Present uint8 //choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	Message NgapMessage
}

//represent InitiatingMessage, SuccessfulOutcome or UnsuccessfulOutcome
type NgapMessage struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality
	Msg           MessageUnmarshaller //to be decoded message
}

//interface to message decoder
//all message need to implement this interface
type MessageUnmarshaller interface {
	decode([]byte) (error, []IeDiagnosticsItem)
}

//API for decode NgapPdu
func NgapDecode(wire []byte) (pdu NgapPdu, err error, diagnostics *IeDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	//1. decode extention bit
	if _, err := r.ReadBit(); err != nil { //extenstion bit is useless for now
		return
	}

	//2. decode present
	present := uint8(0) //TODO: decode
	//3. decode procedure code
	procedureCode := ProcerdureCode(0) //TODO: decode
	//4. decode criticality
	var criticality Criticality //TODO: decode
	//5. decode message content
	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}

	//prepare message for decoding
	message := createMessage(present, procedureCode)
	if message == nil {
		err = fmt.Errorf("Unknown message") //TODO: create a right error message
		return
	}

	var diagnosticsItems []IeDiagnosticsItem
	//decode all IEs within the message
	if err, diagnosticsItems = message.decode(containerBytes); err != nil {
		return
	}
	pdu = NgapPdu{
		Present: present,
		Message: NgapMessage{
			ProcedureCode: procedureCode,
			Criticality:   criticality,
			Msg:           msg,
		},
	}

	//in case there was any critical diagnostics, create a report
	diagnostics = buildDiagnostics(present, procedureCode, criticality, diagnosticsItems)
	return
}

//create a message from Present value and ProcedureCode value to prepare for
//decoding
func createMessage(present uint8, procedureCode ProcedureCode) MessageUnmarshaller {
	switch present {
	case NgapPduInitiatingMessage:
		switch procedureCode {
		case NgapProcedureNGSetup:
			return new(NGSetupRequest)

		}

	case NgapPduSuccessfulOutcome:
		switch procedureCode {
		case NgapProcedureNGSetup:
			return new(NGSetupResponse)
		}

	case NgapPduUnsuccessfulOutcome:
		switch procedureCode {
		case NgapProcedureNGSetup:
			return new(NGSetupFailure)
		}
	}
	return nil
}

func buildDiagnostics(present uint8, procedureCode ProcedureCode, criticality Criticality, diagnosticsItems []DiagnosticsItem) (diagnostics *IeDiagnostics) {
	//TODO: build diagnostic content
	return
}
