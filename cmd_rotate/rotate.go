package cmd_rotate

import (
	"crypto/rsa"
	"fmt"

	klog "github.com/kmesiab/go-klogger"
	"github.com/spf13/cobra"

	"github.com/kmesiab/go-key-rotator-cli/app"
	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/aws"
	"github.com/kmesiab/go-key-rotator-cli/filesystem"
	"github.com/kmesiab/go-key-rotator-cli/types"
)

var (
	err        error
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type RotateCommand struct {
	app.Command
}

func (app RotateCommand) Run(cmd *cobra.Command, _ []string) {
	klog.Logf("Rotating new keys...").Info()

	if !aws.IsValidParameterStoreName(args.GetName(cmd)) {
		klog.Logf(aws.ParameterStoreNamingRequirementsString, args.GetName(cmd)).Error()

		return
	}

	pubKeyName := aws.MakePublicKeyName(args.GetName(cmd))
	privKeyName := aws.MakePrivateKeyName(args.GetName(cmd))

	// Generate and rotate the keys
	privateKey, publicKey, err = app.KeyRotator.RotatePrivateKeyAndPublicKey(
		pubKeyName, privKeyName, app.AWSSession,
	)

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
