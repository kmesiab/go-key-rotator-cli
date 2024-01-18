package types

import (
	"crypto/rsa"

	"github.com/aws/aws-sdk-go/aws/session"
	rotator "github.com/kmesiab/go-key-rotator"
)

type KeyRotatorInterface interface {
	RotatePrivateKeyAndPublicKey(
		string, string, *session.Session,
	) (*rsa.PrivateKey, *rsa.PublicKey, error)
	EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte
	EncodePublicKeyToPEM(publicKey *rsa.PublicKey) ([]byte, error)
}

type KeyRotator struct{}

func (k KeyRotator) RotatePrivateKeyAndPublicKey(
	privateKeyName, publicKeyName string, sess *session.Session,
) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	return rotator.RotatePrivateKeyAndPublicKey(privateKeyName, publicKeyName, sess)
}

func (k KeyRotator) EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	return rotator.EncodePrivateKeyToPEM(privateKey)
}

func (k KeyRotator) EncodePublicKeyToPEM(publicKey *rsa.PublicKey) ([]byte, error) {
	return rotator.EncodePublicKeyToPEM(publicKey)
}
