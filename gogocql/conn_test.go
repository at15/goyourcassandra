package gogocql

import (
	"testing"

	requir "github.com/stretchr/testify/require"
)

func TestConn_Dial(t *testing.T)  {
	require := requir.New(t)
	conn := Conn{}
	err := conn.dial()
	require.Nil(err)
	err = conn.options()
	t.Log(err)
}