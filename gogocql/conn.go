package gogocql

import (
	"net"

	"github.com/dyweb/gommon/errors" // TODO: define a better error message format and contribute back to upstream
	"fmt"
	"encoding/binary"
)

//type ConnConfig struct {
//	HostPort string
//}

type Conn struct {
	//cfg  ConnConfig

	conn net.Conn
}

func (c *Conn) dial() error {
	addr := "localhost:9042"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return errors.Wrapf(err, "dial %s failed", addr)
	}
	c.conn = conn
	// TODO: how to know the connection is broken? write something?
	return nil
}

//     0         8        16        24        32         40
//    +---------+---------+---------+---------+---------+
//    | version |  flags  |      stream       | opcode  |
//    +---------+---------+---------+---------+---------+
//    |                length                 |
//    +---------+---------+---------+---------+
func (c *Conn) options() error {
	var reqV4 byte
	reqV4 = 0x04 // 0 for request, 4 for version
	//var streamId int16
	//streamId = 1 // if we are sync, we don't need streamId
	var opCode byte
	opCode = 0x05 // OPTIONS
	frame := make([]byte, 5+4)
	frame[0] = reqV4 // version
	frame[1] = 0     // flags
	frame[2] = 0     // streamId // TODO: it's bigendian so ...?
	frame[3] = 1     // streamId
	frame[4] = opCode
	_, err := c.conn.Write(frame)
	if err != nil {
		return errors.Wrap(err, "error write frame")
	}
	// TODO: what size of buf should I have
	// TODO: what's the behavior of read ... how can I know if there is more data left?
	var buf = make([]byte, 1024)
	n, err := c.conn.Read(buf)
	if err != nil {
		return errors.Wrap(err, "error read frame")
	}
	fmt.Printf("len %d val %v\n", n, buf)
	// TODO: why first byte is 132? ... oops, it's version I guess? 0x84 -> 8 * 16 + 4 = 132 ... ok
	// [132 0 0 1 6 0 0 0 52 0 2 0 11 67 81 76 95 86 69 82 83 73 79 78 0 1 0 5 51 46 51 46 49 0 11 67 79 77 80 82 69 83 83 73 79 78 0 2 0 6 115 110 97 112 112 121 0 3 108 122 52 0 0 0 0]
	// header 5 bytes, length 4 bytes, body ...
	resOpCode := buf[4]
	fmt.Printf("res op %X\n", resOpCode) // should be 0x06 supported, // TODO: print as 0x06 instead of just 6
	bLen := binary.BigEndian.Uint32(buf[5 : 5+4])
	fmt.Printf("res length %d\n", bLen)
	// options response is [string multimap]
	// `[short]` n, followed by n pair `<k><v>` where `<v>` is `[string list]`
	//nPairs := int16(buf[5+4]<<8 + buf[5+5]) // TODO: what is the right way to read 2 bytes to int16
	nPairs := int16(buf[5+4])<<8 + int16(buf[5+5]) // TODO: is | faster than + ?
	fmt.Printf("res pairs %d\n", nPairs)
	offset := 5 + 4 + 2// header + length + nPairs
	options := make(map[string][]string)
	for i := 0; i < int(nPairs); i++ {
		// read <k><v>
		// k is [string]: [short] n, n bytes
		ksLen := int16(buf[offset])<<8 + int16(buf[offset+1])
		fmt.Printf("i %d ksLen %d\n", i, ksLen)
		offset += 2
		ks := buf[offset : offset+int(ksLen)]
		fmt.Printf("ks %s\n", ks)
		offset += int(ksLen)
		// v is [string list]: [short] n, n strings
		nStrings := int16(buf[offset])<<8 + int16(buf[offset+1])
		offset += 2
		vList := make([]string, 0, nStrings)
		for j := 0; j < int(nStrings); j++ {
			vsLen := int16(buf[offset])<<8 + int16(buf[offset+1])
			fmt.Printf("i %d j %d vsLen %d\n", i, j, vsLen)
			offset += 2
			vs := string(buf[offset : offset+int(vsLen)])
			offset += int(vsLen)
			vList = append(vList, vs)
		}
		options[string(ks)] = vList
	}
	fmt.Printf("options %v\n", options)
	// FIXME: it seems I am not reading string properly ...
	// options map[CQL_VERSION  3.3.1 COMPRESSION  snappy lz4                   :[] :[]]--- PASS: TestConn_Dial (0.00s)
	return nil
}
