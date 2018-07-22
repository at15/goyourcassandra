package frame

// Header is the 9 bytes frame header
type Header struct {
	version  Version
	flags    byte
	streamId int16
	opCode   Op
	length   int32
}

// Body is the frame body with length defined in frame header
type Body struct {

}