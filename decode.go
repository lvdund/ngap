package ngap

import (
	"bytes"
	"fmt"
	"ngap/aper"
	"ngap/ie"
)

// hold a decoded Ngap message
type NgapPdu struct {
	Present uint8 //choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	Message NgapMessage
}

// represent InitiatingMessage, SuccessfulOutcome or UnsuccessfulOutcome
type NgapMessage struct {
	ProcedureCode ProcedureCode
	Criticality   ie.Criticality
	Msg           MessageUnmarshaller //to be decoded message
}

// interface to message decoder
// all message need to implement this interface
type MessageUnmarshaller interface {
	decode([]byte) (error, []ie.CriticalityDiagnostics)
	Encode(aper.AperWriter) error
}

// API for decode NgapPdu
func NgapDecode(wire []byte) (pdu NgapPdu, err error, diagnostics *ie.CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	//1. decode extention bit
	// if _, err := r.ReadBit(); err != nil { //extenstion bit is useless for now
	// 	return
	// }

	//2. decode present		//choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	v, err := r.ReadInteger(&aper.Constrain{Lb: 0, Ub: 2}, true)
	if err != nil {
		return
	}
	var present uint8 = uint8(v)
	//3. decode procedure code
	v, err = r.ReadInteger(&aper.Constrain{Lb: 0, Ub: 255}, true)
	if err != nil {
		return
	}
	var procedureCode ProcedureCode = ProcedureCode(v)
	//4. decode criticality
	c, err := r.ReadEnumerate(&aper.Constrain{Lb: 0, Ub: 2}, true)
	if err != nil {
		return
	}
	var criticality ie.Criticality = ie.Criticality{Value: aper.Enumerated(c)}
	//5. decode message content
	// var containerBytes []byte
	// if containerBytes, err = r.ReadOpenType(); err != nil {
	// 	return
	// }

	//prepare message for decoding
	message := createMessage(present, procedureCode)
	if message == nil {
		err = fmt.Errorf("Unknown message") //TODO: create a right error message
		return
	}

	var diagnosticsItems []ie.CriticalityDiagnostics
	//decode all IEs within the message
	// if err, diagnosticsItems = message.decode(containerBytes); err != nil {
	// 	return
	// }
	pdu = NgapPdu{
		Present: present,
		Message: NgapMessage{
			ProcedureCode: procedureCode,
			Criticality:   criticality,
			Msg:           message,
		},
	}

	//in case there was any critical diagnostics, create a report
	diagnostics = buildDiagnostics(present, procedureCode, criticality, diagnosticsItems)
	return
}

// create a message from Present value and ProcedureCode value to prepare for
// decoding
func createMessage(present uint8, procedureCode ProcedureCode) MessageUnmarshaller {
	switch present {
	case NgapPduInitiatingMessage:
		switch int64(procedureCode) {
		case ie.ProcedureCodeNGSetup:
			return new(NGSetupRequest)
		}

	case NgapPduSuccessfulOutcome:
		switch int64(procedureCode) {
		case ie.ProcedureCodeNGSetup:
			// return new(NGSetupResponse)
			return nil
		}

	case NgapPduUnsuccessfulOutcome:
		switch int64(procedureCode) {
		case ie.ProcedureCodeNGSetup:
			// return new(NGSetupFailure)
			return nil
		}
	}
	return nil
}

func buildDiagnostics(present uint8, procedureCode ProcedureCode, criticality ie.Criticality, diagnosticsItems []ie.CriticalityDiagnostics) (diagnostics *ie.CriticalityDiagnostics) {
	//TODO: build diagnostic content
	return
}
