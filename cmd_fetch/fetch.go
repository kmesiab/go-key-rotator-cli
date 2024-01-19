package cmd_fetch

import (
	"fmt"

	rotator "github.com/kmesiab/go-key-rotator"
	klog "github.com/kmesiab/go-klogger"

	"github.com/kmesiab/go-key-rotator-cli/app"
	"github.com/kmesiab/go-key-rotator-cli/args"
	"github.com/kmesiab/go-key-rotator-cli/aws"
	"github.com/kmesiab/go-key-rotator-cli/filesystem"
	"github.com/kmesiab/go-key-rotator-cli/types"

	"github.com/spf13/cobra"
)

type FetchCommand struct {
	app.Command
}

func (app FetchCommand) Run(cmd *cobra.Command, _ []string) {
	klog.Logf("Fetching keys!").Info()

	if !aws.IsValidParameterStoreName(args.GetName(cmd)) {
		klog.Logf("Invalid Parameter Store name '%s'. Requirements: %s", args.GetName(cmd),
			aws.ParameterStoreNamingRequirementsString).Error()

		return
	}

	pubKeyName := aws.MakePublicKeyName(args.GetName(cmd))
	privKeyName := aws.MakePrivateKeyName(args.GetName(cmd))

	keyRotator := rotator.NewKeyRotator(rotator.NewAWSParameterStore(app.AWSSession))

	privateKey, err := keyRotator.GetCurrentRSAPrivateKey(privKeyName)
	if err != nil {
		klog.Logf("Failed to fetch private key for '%s'. Ensure the key "+
			"exists and you have the necessary permissions.", privKeyName).Add("error", err).Error()

		return
	}

	publicKey, err := keyRotator.GetCurrentRSAPublicKey(pubKeyName)
	if err != nil {
		klog.Logf("Failed to fetch public key for '%s'. Ensure the key "+
			"exists and you have the necessary permissions.", pubKeyName).Add("error", err).Error()

		return
	}

	rotatorResult := &types.Rotation{
		PublicKey:      publicKey,
		PrivateKey:     privateKey,
		PublicKeyName:  aws.GetFilenameFromParameterStorePath(pubKeyName),
		PrivateKeyName: aws.GetFilenameFromParameterStorePath(privKeyName),
	}

	err = filesystem.WriteAllKeysToFile(rotatorResult, keyRotator)
	if err != nil {
		klog.Logf("Failed to write keys to file for '%s': %s. Check file permissions and "+
			"availability of file system.", args.GetName(cmd), err).Error()
		return
	}

	fmt.Printf(`
🔐 Downoaded RSA key pair with names:
	
   💾 Public Key: %s
   💾 Private Key: %s
`,
		pubKeyName,
		privKeyName,
	)
}
