package types

import "crypto/rsa"

// Rotation holds an RSA key pair and their corresponding names. It is used for operations
// involving key generation, rotation, and retrieval. PublicKey and PrivateKey are the RSA
// keys, while PublicKeyName and PrivateKeyName are their respective identifiers, useful
// for storage and retrieval in systems like AWS Parameter Store.
type Rotation struct {
	PublicKey      *rsa.PublicKey
	PrivateKey     *rsa.PrivateKey
	PublicKeyName  string
	PrivateKeyName string
}
