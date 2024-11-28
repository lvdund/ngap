package ngap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

var inituemsg []byte = []byte{0, 15, 64, 66, 0, 0, 5, 0, 85, 0, 2, 0, 1, 0, 38, 0, 24, 23, 126, 0, 65, 121, 0, 13, 1, 2, 248, 57, 0, 0, 0, 0, 0, 0, 0, 0, 16, 46, 2, 160, 32, 0, 121, 0, 15, 64, 2, 248, 57, 0, 0, 8, 0, 0, 2, 248, 57, 0, 0, 1, 0, 90, 64, 1, 24, 0, 112, 64, 1, 0}

func Test2(t *testing.T) {
	// test decode
	if decode, err, _ := NgapDecode(bytes.NewBuffer(inituemsg)); err != nil {
		fmt.Println(err)
	} else {
		msg := decode.Message.Msg.(*ies.InitialUEMessage)
		fmt.Println(msg.RANUENGAPID.Value)
		fmt.Println(msg.NASPDU.Value)
		fmt.Println(msg.UserLocationInformation.Choice)
	}
	// test encode
	msg := GetInitialUEMessage()
	if b, err := NgapEncode(&msg); err != nil {
		fmt.Println("err encode:", err.Error())
	} else {
		fmt.Println(b)
		fmt.Printf("%.8b",b)
	}
}

func GetInitialUEMessage() ies.InitialUEMessage {
	msg := ies.InitialUEMessage{}

	msg.RANUENGAPID = &ies.RANUENGAPID{
		Value: 1,
	}

	msg.NASPDU = &ies.NASPDU{Value: aper.OctetString{126, 0, 65, 121, 0, 13, 1, 2, 248, 57, 0, 0, 0, 0, 0, 0, 0, 0, 16, 46, 2, 160, 32}}

	msg.UserLocationInformation = &ies.UserLocationInformation{
		Choice: ies.UserLocationInformationPresentUserLocationInformationNR,
		UserLocationInformationNR: &ies.UserLocationInformationNR{
			NRCGI: &ies.NRCGI{
				PLMNIdentity: &ies.PLMNIdentity{Value: aper.OctetString{2, 248, 57}},
				NRCellIdentity: &ies.NRCellIdentity{
					Value: aper.BitString{
						Bytes:   []byte{0, 0, 8, 0, 0},
						NumBits: 36,
					},
				},
			},
			TAI: &ies.TAI{
				PLMNIdentity: &ies.PLMNIdentity{Value: aper.OctetString{2, 248, 57}},
				TAC:          &ies.TAC{Value: aper.OctetString{0, 0, 1}},
			},
		},
	}

	msg.RRCEstablishmentCause = &ies.RRCEstablishmentCause{Value: ies.RRCEstablishmentCauseMosignalling}
	msg.UEContextRequest = &ies.UEContextRequest{Value: 0}
	return msg
}

var DownlinkNASTransport1 []byte = []byte{0, 4, 64, 62, 0, 0, 3, 0, 10, 0, 2, 0, 4, 0, 85, 0, 2, 0, 1, 0, 38, 0, 43, 42, 126, 0, 86, 0, 2, 0, 0, 33, 39, 153, 5, 44, 165, 108, 155, 49, 159, 167, 131, 170, 13, 238, 145, 136, 32, 16, 182, 175, 242, 250, 217, 16, 128, 0, 171, 86, 188, 82, 80, 119, 243, 189}

func Test1(t *testing.T) {
	// test decode
	if decode, err, _ := NgapDecode(bytes.NewBuffer(DownlinkNASTransport1)); err != nil {
		fmt.Println(err)
	} else {
		msg := decode.Message.Msg.(*ies.DownlinkNASTransport)
		fmt.Println(msg.AMFUENGAPID.Value)
		fmt.Println(msg.RANUENGAPID.Value)
		fmt.Println(msg.NASPDU.Value)
	}
	// test encode
	msg := BuildDownlinkNASTransport()
	if b, err := NgapEncode(&msg); err != nil {
		fmt.Println("err encode:", err.Error())
	} else {
		fmt.Println("enode:", b)
	}
}

func BuildDownlinkNASTransport() (msg ies.DownlinkNASTransport) {
	nasPdu := []byte{126, 0, 86, 0, 2, 0, 0, 33, 39, 153, 5, 44, 165, 108, 155, 49, 159, 167, 131, 170, 13, 238, 145, 136, 32, 16, 182, 175, 242, 250, 217, 16, 128, 0, 171, 86, 188, 82, 80, 119, 243, 189}
	msg = ies.DownlinkNASTransport{
		AMFUENGAPID: &ies.AMFUENGAPID{Value: 4},
		RANUENGAPID: &ies.RANUENGAPID{Value: 1},
		NASPDU: &ies.NASPDU{
			Value: nasPdu,
		},
	}
	return
}
