package gogocql

import (
	"testing"

	requir "github.com/stretchr/testify/require"
)

func TestConn_Dial(t *testing.T) {
	t.Skip("require cassandra running")

	require := requir.New(t)
	conn := Conn{}
	err := conn.dial()
	require.Nil(err)
	err = conn.options()
	t.Log(err)
}
