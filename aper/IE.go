package aper

type IE interface {
	Encode(*AperWriter) error
	Decode(*AperReader) error
}