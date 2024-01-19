package cmd_rotate

import (
	"crypto/rsa"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	klog "github.com/kmesiab/go-klogger"
	"github.com/spf13/cobra"

	"github.com/kmesiab/go-key-rotator-cli/app"
	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/aws"
	"github.com/kmesiab/go-key-rotator-cli/filesystem"
	"github.com/kmesiab/go-key-rotator-cli/types"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type RotateCommand struct {
	app.Command

	Session    *session.Session
	KeyRotator types.KeyRotatorInterface
}

func (app RotateCommand) Run(cmd *cobra.Command, _ []string) {
	klog.Logf("Rotating new keys...").Info()

	if !aws.IsValidParameterStoreName(args.GetName(cmd)) {
		klog.Logf(aws.ParameterStoreNamingRequirementsString, args.GetName(cmd)).Error()

		return
	}

	pubKeyName := aws.MakePublicKeyName(args.GetName(cmd))
	privKeyName := aws.MakePrivateKeyName(args.GetName(cmd))

	size := args.GetSize(cmd)
	sizeInt, err := strconv.ParseInt(size, 10, 64)

	if err != nil || sizeInt < 2048 || sizeInt > 4096 {

		klog.Logf("Invalid key size: %s. Key size must "+
			"be between %d and %d bits.", size, 2048, 4096).Error()

		return
	}

	// Generate and rotate the keys
	privateKey, publicKey, err = app.KeyRotator.Rotate(pubKeyName, privKeyName, int(sizeInt))

	if err != nil {
		klog.Logf("Error rotating keys: %s\n", err).Error()

		return
	}

	rotationResult := &types.Rotation{
		PublicKey:      publicKey,
		PrivateKey:     privateKey,
		PublicKeyName:  pubKeyName,
		PrivateKeyName: privKeyName,
	}

	// Save the keys to disk
	if err = filesystem.WriteAllKeysToFile(rotationResult, app.KeyRotator); err != nil {
		klog.Logf("Error saving keys to disk: %s\n", err).Error()
	}

	fmt.Printf(`
🔐 Generated and stored keys:
	
   💾 Public Key: %s
   💾 Private Key: %s
`,
		pubKeyName,
		privKeyName,
	)
}
