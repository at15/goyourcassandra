package gogocql

import "testing"

// test binary operations

func TestReadShort(t *testing.T) {
	s := []byte{0, 1}
	i := int16(s[0]<<8 + s[1])
	i2 := int16(s[0])<<8 + int16(s[1])
	t.Logf("i %d i2 %d", i, i2)
	s = []byte{1, 1} // 0000 0001 0000 0001 = 256 + 1
	i = int16(s[0]<<8 + s[1])
	i2 = int16(s[0])<<8 + int16(s[1])
	t.Logf("i %d i2 %d", i, i2)

	// TODO: test if | is faster than +
}
