package filesystem

import (
	"fmt"
	"os"

	rotator "github.com/kmesiab/go-key-rotator"

	"github.com/kmesiab/go-key-rotator-cli/aws"
	"github.com/kmesiab/go-key-rotator-cli/types"
)

func WriteAllKeysToFile(result *types.Rotation, keyRotator types.KeyRotatorInterface) error {
	var (
		err              error
		encodedPublicKey []byte
	)

	privKeyFileName := aws.GetFilenameFromParameterStorePath(result.PrivateKeyName)
	pubKeyFileName := aws.GetFilenameFromParameterStorePath(result.PublicKeyName)

	if err := WritePEMToFile(privKeyFileName, rotator.EncodePrivateKeyToPEM(result.PrivateKey)); err != nil {
		return fmt.Errorf("error writing private key to file: %s\n", err)
	}

	if encodedPublicKey, err = rotator.EncodePublicKeyToPEM(result.PublicKey); err != nil {
		return fmt.Errorf("error encoding public key: %s\n", err)
	}

	if err := WritePEMToFile(pubKeyFileName, encodedPublicKey); err != nil {
		return fmt.Errorf("error writing private key to file: %s\n", err)
	}

	return nil
}

func WritePEMToFile(fileName string, pemData []byte) error {
	return os.WriteFile(fileName, pemData, 0o644)
}
