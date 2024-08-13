package ecies

import (
	"crypto/elliptic"
	"fmt"
)

func example1() {
	curve := elliptic.P256()

	k, err := GenerateKey(curve)
	if err != nil {
		panic("failed to generate key pair")
	}
	privateKey := k
	publicKey := k.PublicKey

	serializedPrivateKey := HexEncode(SerializePrivateKey(privateKey))
	serializedPublicKey := HexEncode(SerializePublicKey(publicKey))
	fmt.Println("privateKey=" + serializedPrivateKey)
	fmt.Println("publicKey=" + serializedPublicKey)

	deserializedPrivateKey := DeserializePrivateKey(k.Curve, HexDecodeWithoutError(serializedPrivateKey))
	deserializedPublicKey, err := DeserializePublicKey(k.Curve, HexDecodeWithoutError(serializedPublicKey))
	if err != nil {
		panic("failed to deserialize key pair")
	}
	fmt.Println(deserializedPrivateKey.D.Cmp(privateKey.D) == 0)
	fmt.Println(deserializedPublicKey.X.Cmp(publicKey.X) == 0)
	fmt.Println(deserializedPublicKey.Y.Cmp(publicKey.Y) == 0)

	ecies := NewECIES(curve)
	plainText := "hello world"
	encryptedTextBytes, err := ecies.Encrypt(publicKey, []byte(plainText))
	if err != nil {
		panic("failed to encrypt message")
	}
	decryptedTextBytes, err := ecies.Decrypt(privateKey, encryptedTextBytes)
	if err != nil {
		panic("failed to decrypt message")
	}
	fmt.Println(string(decryptedTextBytes))
}
