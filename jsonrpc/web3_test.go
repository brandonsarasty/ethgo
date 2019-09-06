package jsonrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
)

func TestWeb3ClientVersion(t *testing.T) {
	s := NewTestServer(t, nil)
	defer s.Close()

	_, err := s.Client().Web3().ClientVersion()
	assert.NoError(t, err)
}

func TestWeb3Sha3(t *testing.T) {
	s := NewTestServer(t, nil)
	defer s.Close()

	src := []byte{0x1, 0x2, 0x3}

	found, err := s.Client().Web3().Sha3(src)
	assert.NoError(t, err)

	k := sha3.NewLegacyKeccak256()
	k.Write(src)
	expected := k.Sum(nil)

	assert.Equal(t, expected, found)
}
