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

// TODO: the benchmark result is 0 for both of them ....
func BenchmarkReadShort(b *testing.B) {
	buf := []byte{0, 1}
	var iPlus int16
	var iOr int16
	b.Run("use +", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			iPlus = int16(buf[0])<<8 + int16(buf[1])
		}

	})
	b.Run("use |", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			iOr = int16(buf[0])<<8 | int16(buf[1])
		}
	})
	b.Logf("iPlus %d iOr %d", iPlus, iOr)
}
