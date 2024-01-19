package cmd_generate

import (
	"fmt"
	"strconv"

	rotator "github.com/kmesiab/go-key-rotator"
	klog "github.com/kmesiab/go-klogger"
	"github.com/spf13/cobra"

	"github.com/kmesiab/go-key-rotator-cli/app"
	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/aws"
	"github.com/kmesiab/go-key-rotator-cli/filesystem"
)

type GenerateCommand struct {
	app.Command
}

func (app GenerateCommand) Run(cmd *cobra.Command, _ []string) {

	if !aws.IsValidParameterStoreName(args.GetName(cmd)) {
		klog.Logf(aws.ParameterStoreNamingRequirementsString, args.GetName(cmd)).Error()

		return
	}

	klog.Logf("Generating new keys! ").Info()

	keyRotator := rotator.NewKeyRotator(
		rotator.NewAWSParameterStore(app.AWSSession),
	)

	size := args.GetSize(cmd)
	sizeInt, err := strconv.ParseInt(size, 10, 64)

	if err != nil || sizeInt < 2048 || sizeInt > 4096 {

		klog.Logf("Invalid key size: %s. Key size must "+
			"be between %d and %d bits.", size, 2048, 4096).Error()

		return
	}

	publicKey, privateKey, err := keyRotator.GenerateKeyPair(int(sizeInt))

	if err != nil {
		klog.Logf("Failed to generate RSA key pair with size %s bits: %d\n",
			size, err).Error()

		return
	}

	publicKeyPEMBytes, err := rotator.EncodePublicKeyToPEM(publicKey)
	privateKeyPEMBytes := rotator.EncodePrivateKeyToPEM(privateKey)

	if err != nil {
		klog.Logf("Failed to encode RSA public key to PEM format: %s\n", err).Error()

		return
	}

	pubKeyName := aws.MakePublicKeyName(args.GetName(cmd))
	privKeyName := aws.MakePrivateKeyName(args.GetName(cmd))

	if err := filesystem.WritePEMToFile(pubKeyName, publicKeyPEMBytes); err != nil {
		klog.Logf("Failed to write public key to file: %s\n", pubKeyName).Error()

		return
	}

	if err := filesystem.WritePEMToFile(privKeyName, privateKeyPEMBytes); err != nil {
		klog.Logf("Failed to write private key to file: %s\n", privKeyName).Error()

		return
	}

	fmt.Printf(`
🔐 Generated %d bit RSA key pair with names:
	
   💾 Public Key: %s
   💾 Private Key: %s
`,
		sizeInt,
		pubKeyName,
		privKeyName,
	)
}
