package ngap

import (
	"fmt"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

var InitialContextSetup []byte = []byte{0, 14, 0, 128, 160, 0, 0, 9, 0, 10, 0, 2, 0, 3, 0, 85, 0, 2, 0, 1, 0, 28, 0, 7, 0, 2, 248, 57, 202, 254, 0, 0, 0, 0, 5, 2, 1, 1, 2, 3, 0, 119, 0, 9, 8, 0, 4, 0, 0, 0, 0, 0, 0, 0, 94, 0, 32, 206, 95, 98, 156, 163, 174, 116, 221, 109, 37, 95, 134, 126, 99, 8, 187, 95, 19, 3, 95, 23, 52, 107, 185, 18, 252, 178, 206, 240, 225, 31, 23, 0, 36, 64, 4, 0, 2, 248, 57, 0, 34, 64, 8, 17, 16, 0, 0, 0, 255, 255, 0, 0, 38, 64, 52, 51, 126, 2, 144, 13, 43, 124, 1, 126, 0, 66, 1, 1, 119, 0, 11, 242, 2, 248, 57, 202, 254, 0, 0, 0, 0, 3, 84, 7, 0, 2, 248, 57, 0, 0, 1, 21, 5, 4, 1, 1, 2, 3, 33, 1, 0, 94, 1, 6, 22, 1, 44}

func Test5(t *testing.T) {
	if decode, err, _ := NgapDecode(InitialContextSetup); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=======================")
		fmt.Println("=======================")
		fmt.Println("=======================")
		msg := decode.Message.Msg.(*ies.InitialContextSetupRequest)
		fmt.Println(*msg.AMFUENGAPID)
		fmt.Println(*msg.RANUENGAPID)
		fmt.Println(msg.GUAMI.PLMNIdentity, msg.GUAMI.AMFRegionID, msg.GUAMI.AMFSetID, msg.GUAMI.AMFPointer)
		fmt.Println(*msg.AllowedNSSAI.Value[0].SNSSAI.SD, *msg.AllowedNSSAI.Value[0].SNSSAI.SST)
		fmt.Println("UESecurityCapabilities")
		fmt.Println(msg.UESecurityCapabilities.NRencryptionAlgorithms)
		fmt.Println(msg.UESecurityCapabilities.NRintegrityProtectionAlgorithms)
		fmt.Println(msg.UESecurityCapabilities.EUTRAencryptionAlgorithms)
		fmt.Println(msg.UESecurityCapabilities.EUTRAintegrityProtectionAlgorithms)
		fmt.Println("UESecurityCapabilities")
		fmt.Println(*msg.SecurityKey)
		fmt.Println(msg.MobilityRestrictionList.ServingPLMN.Value)
		fmt.Println(msg.MaskedIMEISV.Value)
		fmt.Println(*msg.NASPDU)
	}
}

var UplinkNASTransport []byte = []byte{0, 46, 64, 60, 0, 0, 4, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 1, 0, 38, 0, 22, 21, 126, 0, 87, 45, 16, 65, 78, 141, 241, 6, 191, 206, 138, 133, 88, 37, 195, 14, 23, 250, 76, 0, 121, 64, 15, 64, 2, 248, 57, 0, 0, 8, 0, 0, 2, 248, 57, 0, 0, 1}

func Test3(t *testing.T) {
	// test decode
	if decode, err, _ := NgapDecode(UplinkNASTransport); err != nil {
		fmt.Println(err)
	} else {
		msg := decode.Message.Msg.(*ies.UplinkNASTransport)
		fmt.Println(msg.RANUENGAPID.Value)
		fmt.Println(msg.AMFUENGAPID.Value)
		fmt.Println(msg.NASPDU.Value)
	}
	// test encode
	// msg := GetInitialUEMessage()
	// if b, err := NgapEncode(&msg); err != nil {
	// 	fmt.Println("err encode:", err.Error())
	// } else {
	// 	fmt.Println(b)
	// 	fmt.Printf("%.8b", b)
	// }
}

var inituemsg []byte = []byte{0, 15, 64, 66, 0, 0, 5, 0, 85, 0, 2, 0, 1, 0, 38, 0, 24, 23, 126, 0, 65, 121, 0, 13, 1, 2, 248, 57, 0, 0, 0, 0, 0, 0, 0, 0, 16, 46, 2, 160, 32, 0, 121, 0, 15, 64, 2, 248, 57, 0, 0, 8, 0, 0, 2, 248, 57, 0, 0, 1, 0, 90, 64, 1, 24, 0, 112, 64, 1, 0}

func Test2(t *testing.T) {
	// test decode
	if decode, err, _ := NgapDecode(inituemsg); err != nil {
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
		fmt.Printf("%.8b", b)
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
	if decode, err, _ := NgapDecode(DownlinkNASTransport1); err != nil {
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
