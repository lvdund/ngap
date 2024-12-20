<br />
<div align="center">
  <h3 align="center">Next generation application protocol (NGAP)</h3>
  <p align="center">
    Publish library
  </p>
</div>

## Usage


## TODO list:
 - Encoding API: packet, err := ngap.Encode(msg) -> packet, err := msg.Encode()
 - Should we need a 'ies' submodule or we can move all messages and IEs to the root folder so that we don't need to import multiple submodules? It is the best that users only need to import 'github.com/lvdund/ngap'
 - When composing a message or an IEs, it is best if users do not need to use aper sub-module, except for the case of BitString.
 - IE structure with a single field 'Value' can be replace by type of the Value itself. For example

 ```go
type AMFUENGAPID struct {
	Value aper.Integer
}

type AnNgapMessage struct {
	AMFUENGAPID AMFUENGAPID
	// other fields
}
```

The message should become:

```go
type AnNgapMessage struct {
	AMFUENGAPID integer
}
```	

## Contributing

- [lvdund](https://github.com/lvdund)
- [nguyenducc](https://github.com/nguyenducc)
- [reogac](https://github.com/reogac)
