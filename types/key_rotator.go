package types

import "crypto/rsa"

type KeyRotatorInterface interface {
	Rotate(
		parameterStoreKeyNamePrivateKey,
		parameterStoreKeyNamePublicKey string,
		keySize int,
	) (*rsa.PrivateKey, *rsa.PublicKey, error)
}
