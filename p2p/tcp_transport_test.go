package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {

	tcpOpts := TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder: DefaultDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)

	assert.Equal(t, tr.ListenAddr, ":3000")

	assert.Nil(t, tr.ListenAndAccept())
}