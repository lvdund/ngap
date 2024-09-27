package ngap

import (
	"bytes"
	"fmt"
	"ngap/aper"
	"ngap/ie"

	"github.com/sirupsen/logrus"
)

// hold a decoded Ngap message
type NgapPdu struct {
	Present uint8 //choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	Message NgapMessage
}

// represent InitiatingMessage, SuccessfulOutcome or UnsuccessfulOutcome
type NgapMessage struct {
	ProcedureCode ie.ProcedureCode
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
	fmt.Printf("pdu=%.8b\n", wire)
	r := aper.NewReader(bytes.NewReader(wire))
	//1. decode extention bit
	var b bool
	if b, err = r.ReadBool(); err != nil {
		return
	}
	fmt.Println("extenstion bit:", b)
	//2. decode present		//choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	// v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false)
	c, err := r.ReadChoice(2, false)
	if err != nil {
		return
	}
	present := uint8(c)
	fmt.Println("present:", present)
	//3. decode procedure code
	v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	fmt.Printf("procedureCode:%d-%.8b\n", v, v)
	var procedureCode ie.ProcedureCode = ie.ProcedureCode{Value: aper.Integer(v)}
	//4. decode criticality
	e, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	var criticality ie.Criticality = ie.Criticality{Value: aper.Enumerated(e)}
	fmt.Println("criticality:", e)
	//5. decode message content
	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}
	fmt.Printf("exten=%v - containerBytes readopen type= %.8b\n", b, containerBytes)

	//prepare message for decoding
	message := createMessage(present, procedureCode)
	if message == nil {
		err = fmt.Errorf("Unknown message") //TODO: create a right error message
		return
	}

	var diagnosticsItems []ie.CriticalityDiagnostics
	//decode all IEs within the message
	if err, diagnosticsItems = message.decode(containerBytes); err != nil {
		return
	}
	logrus.Infoln("message body:", message)
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
func createMessage(present uint8, procedureCode ie.ProcedureCode) MessageUnmarshaller {
	switch present {
	case NgapPduInitiatingMessage:
		switch int64(procedureCode.Value) {
		case ie.ProcedureCodeNGSetup:
			return new(NGSetupRequest)
		}

	case NgapPduSuccessfulOutcome:
		switch int64(procedureCode.Value) {
		case ie.ProcedureCodeNGSetup:
			// return new(NGSetupResponse)
			return nil
		}

	case NgapPduUnsuccessfulOutcome:
		switch int64(procedureCode.Value) {
		case ie.ProcedureCodeNGSetup:
			// return new(NGSetupFailure)
			return nil
		}
	}
	return nil
}

func buildDiagnostics(present uint8, procedureCode ie.ProcedureCode, criticality ie.Criticality, diagnosticsItems []ie.CriticalityDiagnostics) (diagnostics *ie.CriticalityDiagnostics) {
	//TODO: build diagnostic content
	return
}
