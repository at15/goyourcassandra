package frame

// Version contains both the direction and proto version, MSB 0 is request, 1 is response
type Version byte

const (
	V4Request  Version = 0x04
	V4Response Version = 0x84
	V5Request  Version = 0x05
	V5Response Version = 0x85
)

// Op defines type of the frame
type Op byte

const (
	OpError        Op = 0x00
	OpStartup      Op = 0x01
	OpReady        Op = 0x02
	OpAuthenticate Op = 0x03
	// TODO: why 0x04 is left blank ... it is not said in the proto doc
	OpOptions       Op = 0x05
	OpSupported     Op = 0x06
	OpQuery         Op = 0x07
	OpResult        Op = 0x08
	OpPrepare       Op = 0x09
	OpExecute       Op = 0x0A
	OpRegister      Op = 0x0B
	OpEvent         Op = 0x0C
	OpBatch         Op = 0x0D
	OpAuthChallenge Op = 0x0E
	OpAuthResponse  Op = 0x0F
	OpAuthSuccess   Op = 0x10
)

// TODO: flags, how to set them via mask
