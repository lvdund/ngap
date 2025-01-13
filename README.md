<br />
<div align="center">
  <h3 align="center">Next generation application protocol (NGAP)</h3>
  <p align="center">
    Publish library
  </p>
</div>

## Usage

- NGAP msg:
```go
msg := ies.NGSetupRequest{    // need check all mandatory fields
  GlobalRANNodeID: ...
  SupportedTAList: ...
  DefaultPagingDRX: ...
}

// encode
var b []byte
var err error
b, err = NgapEncode(&msg)

// decode
var pdu NgapPdu
var cridia *ies.CriticalityDiagnostics
pdu, err, cridia = NgapDecode(b)
decode_msg := pdu.Message.Msg.(*ies.NGSetupRequest)
```

- Transfer IE:
```go
msg := ies.PDUSessionResourceSetupResponseTransfer{
  // check mandatory fields
}

// encode
var b []byte
var err error
b, err = msg.Encode()

// decode
var decode_transfer ies.PDUSessionResourceSetupResponseTransfer{}
err = decode_transfer.Decode(b)
```

## Contributing

- [lvdund](https://github.com/lvdund)
- [nguyenducc](https://github.com/nguyenducc)
- [reogac](https://github.com/reogac)
