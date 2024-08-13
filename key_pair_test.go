package ecies

import (
	"crypto/elliptic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPair(t *testing.T) {
}

func TestGenerateKey(t *testing.T) {
	curve := elliptic.P256()
	{
		privateKey, err := GenerateKey(curve)
		assert.Nil(t, err)
		assert.NotNil(t, privateKey)
		assert.NotNil(t, privateKey.PublicKey)
		assert.NotNil(t, privateKey.D)
		assert.NotNil(t, privateKey.PublicKey.Curve)
		assert.NotNil(t, privateKey.PublicKey.X)
		assert.NotNil(t, privateKey.PublicKey.Y)
		assert.Equal(t, curve, privateKey.Curve)
	}
	{
		k, _ := GenerateKey(curve)
		privateKeyString := HexEncode(SerializePrivateKey(k))
		publicKeyString := HexEncode(SerializePublicKey(k.PublicKey))
		assert.NotNil(t, publicKeyString)
		assert.NotNil(t, privateKeyString)
	}
}

func TestKeySerialization(t *testing.T) {
	curve := elliptic.P256()
	k, err := GenerateKey(curve)
	assert.Nil(t, err)
	{
		publicKeyBytes := SerializePublicKey(k.PublicKey)
		assert.NotNil(t, publicKeyBytes)
		publicKey, err := DeserializePublicKey(curve, publicKeyBytes)
		assert.Nil(t, err)
		assert.NotNil(t, publicKey)
		assert.True(t, bytesEquals(k.X.Bytes(), publicKey.X.Bytes()))
		assert.True(t, bytesEquals(k.Y.Bytes(), publicKey.Y.Bytes()))
	}
	{
		x, y := SerializePublicKeyToCoordinate(k.PublicKey)
		assert.NotNil(t, x)
		assert.NotNil(t, y)
		publicKey, err := DeserializePublicKeyFromCoordinate(curve, x, y)
		assert.Nil(t, err)
		assert.NotNil(t, publicKey)
		assert.True(t, bytesEquals(k.X.Bytes(), publicKey.X.Bytes()))
		assert.True(t, bytesEquals(k.Y.Bytes(), publicKey.Y.Bytes()))
	}
	{
		privateKeyBytes := SerializePrivateKey(k)
		assert.NotNil(t, privateKeyBytes)
		privateKey := DeserializePrivateKey(curve, privateKeyBytes)
		assert.Nil(t, err)
		assert.NotNil(t, privateKey)
		assert.True(t, bytesEquals(k.D.Bytes(), privateKey.D.Bytes()))
	}
}
